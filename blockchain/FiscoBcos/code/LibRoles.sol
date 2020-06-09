pragma solidity ^0.4.4;
//角色管理库
library LibRoles
{
    struct Role{
        mapping (address=>bool) bearer;
    }
    //创建角色
    function add(Role storage role, address account) internal{
        require(!has(role, account), "Roles: account already has role");
        role.bearer[account] = true;
    }
    //删除角色
    function remove(Role storage role, address account) internal{
        require(has(role, account), "Roles: account does not have role");
        role.bearer[account] = false;
    }
    //查询角色
    function has(Role storage role, address account) internal view returns(bool){
        require(account != address(0), "Roles: account is the zero address");
        return role.bearer[account];
    }
    
}