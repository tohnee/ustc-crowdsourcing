pragma solidity ^0.4.4;
pragma experimental ABIEncoderV2;//用于函数返回值为string[]
import "./VersionManager.sol";

//任务存储库
contract TaskRepository is VersionManager
{
	struct Task{
		string taskname;       //任务名
		address owner;         //任务发布者
		uint duration;         //任务持续时间
		uint startTime;        //任务发布时间
		string taskBrief;      //任务简介
		uint expectedPrice;    //预期价格
		bool IsEnd;            //任务是否结束
		bool IsPaid;           //是否给予奖励
	}
	
	mapping(string=>Task) private _Tasks;   //任务集
	string[] private _Tasknames ;           //任务名
	string[] private mytasknames;           //某申请者的任务名

	function setData(string taskname, uint duration, string taskBrief, uint expectedPrice) onlyLatestVersion public{
	    //判断任务是否存在
	    require(!(keccak256(abi.encodePacked(_Tasks[taskname].taskname)) == keccak256(abi.encodePacked(taskname))), "Task already exist");
		_Tasks[taskname].taskname = taskname;
		_Tasks[taskname].owner = tx.origin;
		_Tasks[taskname].duration = duration;
		_Tasks[taskname].startTime = now;
		_Tasks[taskname].taskBrief = taskBrief;
		_Tasks[taskname].expectedPrice = expectedPrice;
		_Tasks[taskname].IsEnd = false;
		_Tasks[taskname].IsPaid = false;
		_Tasknames.push(taskname);
	}

	//根据任务名获取任务数据
	function getData(string taskname) onlyLatestVersion public view returns(string, address, uint, uint, string, uint, bool, bool){
		Task storage tasktemp = _Tasks[taskname];
		////判断任务是否存在   string的比较：根据hash值
		require(keccak256(abi.encodePacked(tasktemp.taskname)) == keccak256(abi.encodePacked(taskname)), "Task not exist");
		return(tasktemp.taskname, tasktemp.owner, tasktemp.duration, tasktemp.startTime, tasktemp.taskBrief, tasktemp.expectedPrice,
		tasktemp.IsEnd, tasktemp.IsPaid);
	}
	
	//获取所有任务名
	function getAllTaskname() onlyLatestVersion public view returns(string[]){
	    return _Tasknames;
	}
	//获取某申请者的所有任务名
	function getAllMyTaskname() onlyLatestVersion  public returns(string[]){
	    //清空mytaskname
	    delete mytasknames;
	    
	     uint lengthtemp = _Tasknames.length;
	     for(uint i=0; i<lengthtemp; ++i){
	         if(_Tasks[_Tasknames[i]].owner == tx.origin)
	            mytasknames.push(_Tasknames[i]);
	     }
	     return mytasknames;
	}
	
	//修改任务状态为已结束
	function taskIsEnd(string taskname) onlyLatestVersion public{
	    _Tasks[taskname].IsEnd = true;
	}
	//修改任务状态为已支付
	function taskIsPaid(string taskname) onlyLatestVersion public{
	    _Tasks[taskname].IsPaid = true;
	}
}