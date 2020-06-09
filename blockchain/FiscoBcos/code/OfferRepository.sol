pragma solidity ^0.4.4;
pragma experimental ABIEncoderV2;
import "./VersionManager.sol";
import "./LibSafeMath.sol";

//报价存储库
contract OfferRepository is VersionManager
{
    using LibSafeMath for uint256;
    
	struct Offer{
		uint arrivalTime;    //到达时间
		uint departureTime;  //离开时间
		uint cost;           //报价
		address owner;       //工人名
		string taskname;     //任务名
	}

	//一个任务名对应多个报价 taskname => offer结构体
	//一个工人对应多个任务名
	mapping(string=>Offer[]) private _TaskOffers;
	mapping(address=>string[]) private _workerTasknames;
	
	function setData(string taskname, uint duration, uint cost) onlyLatestVersion public{
	    //判断报价是否已存在
	    bool Isexist = false;
 		uint lengthone = _workerTasknames[tx.origin].length;
 		for(uint i=0;i<lengthone;++i){
 		    if(keccak256(abi.encodePacked(_workerTasknames[tx.origin][i])) == keccak256(abi.encodePacked(taskname)))
 		        Isexist = true;
 		}
 		require(!Isexist,"Offer already exist");
	    
	    //memory    用tx.origin,不能用msg.sender
		Offer memory offertemp = Offer(now, now.add(duration*3600), cost, tx.origin, taskname);
		_TaskOffers[taskname].push(offertemp);
		_workerTasknames[tx.origin].push(taskname);
	}

	//工人只能查询自己的报价
	function getData(string taskname) onlyLatestVersion public view returns(uint, uint, uint, address, string){
 		//判断报价是否已存在
 		bool Isexist = false;
 		uint lengthone = _workerTasknames[tx.origin].length;
 		for(uint i=0;i<lengthone;++i){
 		    if(keccak256(abi.encodePacked(_workerTasknames[tx.origin][i])) == keccak256(abi.encodePacked(taskname)))
 		        Isexist = true;
 		}
 		require(Isexist,"Offer does not exist");
 		
 		uint lengthTemp = _TaskOffers[taskname].length;
 		for(i=0; i<lengthTemp; ++i){
 			if(_TaskOffers[taskname][i].owner == tx.origin)
 				return(_TaskOffers[taskname][i].arrivalTime, _TaskOffers[taskname][i].departureTime,
 					_TaskOffers[taskname][i].cost, _TaskOffers[taskname][i].owner,
 					_TaskOffers[taskname][i].taskname);
 		}
	}
	
	//获取本工人所有参与报价的任务
	function  getAllMyofferTotaskname() onlyLatestVersion public view returns(string[]){
	    return _workerTasknames[tx.origin];
	}
	
	//任务分配
	//inout: 任务名,预期价格
    //output: 被分配者，报价
	//分配给出价最低的工人，奖励支付金额为除已分配的工人中的最低报价。
	//若只有一个工人报价，奖励金额为该任务的预期报价。
	function taskAssign(string taskname, uint expectedPrice) onlyLatestVersion public view returns(address, uint){
	    //判断该任务是否有人报价
	    uint lengthTemp = _TaskOffers[taskname].length;
 		require(lengthTemp>0, "No offer!");
 		//只有一个工人报价
 		if(lengthTemp == 1){
 		    return(_TaskOffers[taskname][0].owner, expectedPrice);
 		}
 		uint mincost=0;
 		uint bonus=0;
 		address Assignor;
 		uint AssignCost;
 		if(_TaskOffers[taskname][0].cost > _TaskOffers[taskname][1].cost)
 		    mincost = 1;
 		else bonus =1;
 		
 		for(uint i=2; i<lengthTemp; ++i)
 		{
 		    if(_TaskOffers[taskname][i].cost < _TaskOffers[taskname][mincost].cost)
 		    {
 		        bonus = mincost;
 		        mincost = i;
 		    }
 		    else if(_TaskOffers[taskname][i].cost < _TaskOffers[taskname][bonus].cost)
 		        bonus =i;
 		}
 		Assignor = _TaskOffers[taskname][mincost].owner;
 		AssignCost = _TaskOffers[taskname][bonus].cost;
 		return (Assignor, AssignCost);
	}
	
}