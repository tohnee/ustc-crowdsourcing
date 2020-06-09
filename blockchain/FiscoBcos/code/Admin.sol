pragma solidity ^0.4.4;
import "./CrowdsourceController.sol";

contract Admin is BasicAuth
{
    //合约地址
    address public _taskAddress;
    address public _offerAddress;
    address public _resultAddress;
    address public _applicantAddress;
    address public _workerAddress;
    address public _controllerAddress;
    
    //初始化
    constructor() public{
        //创建数据层合约，三个数据，两个角色
        TaskRepository task = new TaskRepository();
        OfferRepository offer = new OfferRepository();
        AssignResult result = new AssignResult();
        ApplicantRole applicant = new ApplicantRole();
        WorkerRole worker = new WorkerRole();
        _taskAddress = address(task);
        _offerAddress = address(offer);
        _resultAddress = address(result);
        _applicantAddress = address(applicant);
        _workerAddress = address(worker);
        
        //根据数据层合约地址创建控制层合约
        CrowdsourceController controller = new CrowdsourceController(_taskAddress, _offerAddress, _resultAddress, _applicantAddress, _workerAddress);
        _controllerAddress = address(controller);
        
        //控制层合约地址即合约版本
        task.upgradeVersion(_controllerAddress);
        offer.upgradeVersion(_controllerAddress);
        result.upgradeVersion(_controllerAddress);
        applicant.upgradeVersion(_controllerAddress);
        worker.upgradeVersion(_controllerAddress);
        
    }
    
    //升级合约
    function upgradeVersion(address newVersion) public onlyOwner{
        //根据数据层合约地址引入先前的数据层合约，保证数据完整。
        TaskRepository task = TaskRepository(_taskAddress);
        OfferRepository offer = OfferRepository(_offerAddress);
        AssignResult result = AssignResult(_resultAddress);
        ApplicantRole applicant = ApplicantRole(_applicantAddress);
        WorkerRole worker = WorkerRole(_workerAddress);
        
        task.upgradeVersion(newVersion);
        offer.upgradeVersion(newVersion);
        result.upgradeVersion(newVersion);
        applicant.upgradeVersion(newVersion);
        worker.upgradeVersion(newVersion);
    }
    
}
