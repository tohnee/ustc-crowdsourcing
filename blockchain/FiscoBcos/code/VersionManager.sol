pragma solidity ^0.4.4;
import "./BasicAuth.sol";

//版本管理
contract VersionManager is BasicAuth
{
    address _latestVersion;         //最新版本(controller)
    
    modifier onlyLatestVersion(){
        require(msg.sender == _latestVersion, "Not latestVersion");
        _;
    }
    
    //更新版本
    function upgradeVersion(address newVersion) public{
        //对于TaskRepository,offerRepository,AssignResult, 实际上是调用该函数的地址，即Admin合约的地址
        require(msg.sender == _owner, "Not owner");  
        _latestVersion = newVersion;
    }
}