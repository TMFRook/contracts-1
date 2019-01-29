package licence_test

import (
	"context"
	"fmt"
	"os"
	"testing"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/tokencard/contracts/pkg/bindings"
	"github.com/ethereum/go-ethereum/core/types"
	. "github.com/tokencard/contracts/test/shared"
)

var Dao *bindings.Dao
var DaoAddress common.Address

func init() {
	TestRig.AddCoverageForContracts(
		"../../build/licence/combined.json",
		 "../../contracts")
}

func TestLicenceSuite(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Contract Suite")
}

var _ = BeforeEach(func() {
	err := InitializeBackend()
	Expect(err).ToNot(HaveOccurred())

	var tx *types.Transaction

	LicenceAddress, tx, Licence, err = bindings.DeployLicence(BankAccount.TransactOpts(), Backend, Owner.Address(), true, big.NewInt(1), common.HexToAddress("0x0"), TokenHolderAddress)//FIX ME: random should become CryptoFloat contract
	Expect(err).ToNot(HaveOccurred())
	Backend.Commit()
	Expect(isSuccessful(tx)).To(BeTrue())

	DaoAddress, tx, Dao, err = bindings.DeployDao(BankAccount.TransactOpts(), Backend, Owner.Address(), true, LicenceAddress)
	Expect(err).ToNot(HaveOccurred())
	Backend.Commit()
	Expect(isSuccessful(tx)).To(BeTrue())

})

var _ = AfterSuite(func() {
	TestRig.ExpectMinimumCoverage("licence.sol", 0.00)
	TestRig.PrintGasUsage(os.Stdout)
})

func isSuccessful(tx *types.Transaction) bool {
	r, err := Backend.TransactionReceipt(context.Background(), tx.Hash())
	Expect(err).ToNot(HaveOccurred())
	return r.Status == types.ReceiptStatusSuccessful
}

func isGasExhausted(tx *types.Transaction, gasLimit uint64) bool {
	r, err := Backend.TransactionReceipt(context.Background(), tx.Hash())
	Expect(err).ToNot(HaveOccurred())
	if r.Status == types.ReceiptStatusSuccessful {
		return false
	}
	return r.GasUsed == gasLimit
}

var _ = AfterEach(func() {
	td := CurrentGinkgoTestDescription()
	if td.Failed {
		fmt.Fprintf(GinkgoWriter, "\nLast Executed Smart Contract Line for %s:%d\n", td.FileName, td.LineNumber)
		fmt.Fprintln(GinkgoWriter, TestRig.LastExecuted())
	}
})
