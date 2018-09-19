pragma solidity ^0.4.24;


/// @title Controller stores a list of controller addresses that can be used for authentication in other contracts.
contract Controller {
    event AddController(address _sender, address _controller);
    event RemoveController(address _sender, address _controller);

    mapping (address => bool) private _isController;
    uint private _controllerCount;

    /// @dev Checks if message sender is a controller.
    modifier onlyController() {
        require(isController(msg.sender), "sender is not a controller");
        _;
    }

    /// @dev Constructor initializes the list of controllers with the provided address.
    /// @param _account address to add to the list of controllers.
    constructor(address _account) public {
        _addController(_account);
    }

    /// @return true if the provided account is a controller.
    function isController(address _account) public view returns (bool) {
        return _isController[_account];
    }

    /// @return the current number of controllers.
    function controllerCount() public view returns (uint) {
        return _controllerCount;
    }

    /// @dev Add a new controller to the list of controllers.
    /// @param _account address to add to the list of controllers.
    function addController(address _account) external onlyController {
        _addController(_account);
    }

    /// @dev Remove a controller from the list of controllers.
    /// @param _account address to remove from the list of controllers.
    function removeController(address _account) external onlyController {
        _removeController(_account);
    }

    /// @dev Internal-only function that adds a new controller.
    function _addController(address _account) internal {
        require(!_isController[_account], "provided account is already a controller");
        _isController[_account] = true;
        _controllerCount++;
        emit AddController(msg.sender, _account);
    }

    /// @dev Internal-only function that removes an existing controller.
    function _removeController(address _account) internal {
        require(_isController[_account], "provided account is not a controller");
        require(_controllerCount > 1, "cannot remove the last controller");
        _isController[_account] = false;
        _controllerCount--;
        emit RemoveController(msg.sender, _account);
    }
}
