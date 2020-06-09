pragma solidity ^0.4.4;
import "./LibRoles.sol";
import "./LibSafeMath.sol";
import "./VersionManager.sol";

//申请者合约
contract ApplicantRole is VersionManager
{
    using LibSafeMath for uint256;
    using LibRoles for LibRoles.Role;
    LibRoles.Role private _applicants;
    mapping(address => uint) private _balances;  //账号余额
    
    event addApplicantEvent(address indexed account);
    //是否申请者
    function isApplicant(address account) public view returns(bool){
        return _applicants.has(account);
    }
    
    //添加申请者
    function addApplicant(address account) onlyLatestVersion public {
        _applicants.add(account);
        _balances[account] = 1000;
        emit addApplicantEvent(account);
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
}