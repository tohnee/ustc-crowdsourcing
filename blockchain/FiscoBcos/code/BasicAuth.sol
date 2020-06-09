pragma solidity ^0.4.4;
//基础权限合约
contract BasicAuth
{
	address public _owner;

	constructor() public{
		_owner = msg.sender;
	}

	modifier onlyOwner(){
		require(msg.sender == _owner, "Not admin");
		_;
	}
}