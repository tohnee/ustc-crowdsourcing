pragma solidity ^0.4.4;
import "./VersionManager.sol";

//分配结果存储库
contract AssignResult is VersionManager
{
    struct Result{
        string taskname;     //任务名
        address assgignor;   //被分配者
        uint cost;           //奖励支付
    }
    
    //分配结果 结构体数组
    Result[] private _Results;
    
    function setData(string taskname, address assgignor, uint cost) onlyLatestVersion public{
        Result memory resulttemp = Result(taskname, assgignor, cost);
        _Results.push(resulttemp);
    }
    
    //根据任务名获取任务分配结果
    function getData(string taskname) onlyLatestVersion public view returns(string, address, uint){
        //判断该任务是否已分配
        uint lengthTemp = _Results.length;
        uint len =0;
        for(uint i=0; i<lengthTemp; ++i){
            if(keccak256(abi.encodePacked(taskname)) != keccak256(abi.encodePacked(_Results[i].taskname)))
                len++;
        }
        require(len != lengthTemp, "This task does not exist!");
        for(i=0; i<lengthTemp; ++i){
            if(keccak256(abi.encodePacked(taskname)) == keccak256(abi.encodePacked(_Results[i].taskname)))
                return(_Results[i].taskname, _Results[i].assgignor, _Results[i].cost);
        }
    }
}