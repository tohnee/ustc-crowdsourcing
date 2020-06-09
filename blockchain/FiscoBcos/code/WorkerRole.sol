pragma solidity ^0.4.4;
import "./LibRoles.sol";
import "./LibSafeMath.sol";
import "./VersionManager.sol";

//工人合约
contract WorkerRole is VersionManager
{
    using LibSafeMath for uint256;
    using LibRoles for LibRoles.Role;
    LibRoles.Role private _workers;
    mapping(address => uint) private _balances;  //账号余额
    mapping(address => uint) private _creditScore;  //用户信用积分
    
    event AddworkerEvent(address indexed account);
    //是否工人
    function isWorker(address account) public view returns(bool){
        return _workers.has(account);
    }
    
    //添加工人
    function addWorker(address account) onlyLatestVersion public {
        _workers.add(account);
        _balances[account] = 1000;
        _creditScore[account] = 100;
        emit AddworkerEvent(account);
    }
    
    function getBalance(address account) onlyLatestVersion view public returns(uint){
        return _balances[account];
    }
    function addBalance(address account, uint amount) onlyLatestVersion public {
        _balances[account] = _balances[account].add(amount);
    }
    function subBalance(address account, uint amount) onlyLatestVersion public {
        _balances[account] = _balances[account].sub(amount);
    }
    function subCreditScore(address account) onlyLatestVersion public {
        if(_creditScore[account]==0)
            return;
        _creditScore[account] = _creditScore[account].sub(10);
    }
    function getCreditScore(address account) onlyLatestVersion view public returns(uint){
        return _creditScore[account];
    }
}