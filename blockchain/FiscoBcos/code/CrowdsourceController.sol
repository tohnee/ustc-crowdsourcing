pragma solidity ^0.4.4;
pragma experimental ABIEncoderV2;
import "./TaskRepository.sol";
import "./OfferRepository.sol";
import "./AssignResult.sol";
import "./WorkerRole.sol";
import "./ApplicantRole.sol";
contract CrowdsourceController
{
    using LibSafeMath for uint256;
    
    TaskRepository public _task;
    OfferRepository public _offer;
    AssignResult public _result;
    ApplicantRole public _applicant;
    WorkerRole public _worker;
    
    //事件（数据写入相关）
    event ReleaseTaskEvent(string taskname, address owner, uint duration, uint startTime, string taskBrief, uint expectedPrice);
    event WorkerOfferEvent(uint arrivalTime, uint departureTime, uint cost, address owner, string taskname);
    event TaskAssignmentEvent(string taskname, address assignor, uint Cost);
    event BonusPaymentEvent(string taskname, address assignor, uint Cost, uint creditScore, uint applicantBalance, uint workerBalance);
    
    //构造器 根据Admin传入的五个合约地址引入五个数据层合约
    constructor(address _taskAddress, address _offerAddress, address _resultAddress, address _applicantAddress, address _workerAddress) public{
        _task = TaskRepository(_taskAddress);
        _offer = OfferRepository(_offerAddress);
        _result = AssignResult(_resultAddress);
        _applicant = ApplicantRole(_applicantAddress);
        _worker = WorkerRole(_workerAddress);
    }
    //修改器
    modifier onlyApplicant(){
        require(_applicant.isApplicant(msg.sender), "ApplicantRole: caller does not have the role");
        _;
    }
    modifier onlyWorker(){
        require(_worker.isWorker(msg.sender), "WorkerRole: caller does not have the role");
        _;
    }
    
    //函数（数据写入相关）
    //--------1.添加自己为申请者-------------
    function addApplicant() public{
        _applicant.addApplicant(msg.sender);
    }
    
    //--------2.添加自己为工人-------------
    function addWorker() public{
        _worker.addWorker(msg.sender);
    }
    
    //--------3.申请者发布任务------------
    //input: 任务名，任务持续时间
    function releaseTask(string taskname, uint duration, string taskBrief, uint expectedPrice) onlyApplicant public{
        _task.setData(taskname, duration, taskBrief, expectedPrice);
        emit ReleaseTaskEvent(taskname, msg.sender, duration, now, taskBrief, expectedPrice);
    }
    
    //------- 4.工人报价----------------
    //input: 任务名，报价持续时间，报价
    function workerOffer(string taskname, uint duration, uint cost) onlyWorker public{
        //判断任务是否结束
        bool IsEnd;
        (,,,,,,IsEnd,) = _task.getData(taskname);
        require(!IsEnd,"The task has already ended!");
        
        _offer.setData(taskname, duration, cost);
        emit WorkerOfferEvent(now, now.add(duration*3600), cost, msg.sender, taskname);
    }
    
    //-------5.申请者任务分配------------
    //inout: 任务名
    //output: 任务名，被分配者，报价
    function taskAssignment(string taskname) onlyApplicant public returns(string Taskname, address assignor, uint Cost){
        address owner;
        bool IsEnd;
        uint expectedPrice;
        (,owner,,,,expectedPrice,IsEnd,) = _task.getData(taskname);
        //只能分配自己的任务
        require(owner == msg.sender, "Not your task!");
        //判断任务是否已分配
        require(!IsEnd,"The task has been assigned!");
        
        address Assignor;
        uint cost;
        (Assignor,cost) = _offer.taskAssign(taskname,expectedPrice);
        emit TaskAssignmentEvent(taskname, Assignor, cost);
        //修改task状态
        _task.taskIsEnd(taskname);
        _result.setData(taskname, Assignor, cost);
        return(taskname, Assignor, cost);
    }
    
    //-------6.奖励支付-----------
    //input: 任务名，是否合格
    function BonusPayment(string taskname, bool IsQualified) onlyApplicant public{
        address owner;
        bool IsEnd;
        bool IsPaid;
        (,owner,,,,,IsEnd,IsPaid)= _task.getData(taskname);
        //只能对自己的任务进行奖励支付
        require(owner == msg.sender, "Not your task!");
        //判断任务是否已分配且未支付
        require(IsEnd && !IsPaid,"Task has been paid or has not been assigned!");
        
        address assignor;
        uint cost;
        (,assignor,cost) = _result.getData(taskname);
        //修改信用分，奖励支付
        if(IsQualified == false)
            _worker.subCreditScore(assignor);
            
        cost = (cost * _worker.getCreditScore(assignor))/100;
        _applicant.subBalance(owner, cost);
        _worker.addBalance(assignor, cost);
        
        //修改task状态
        _task.taskIsPaid(taskname);
        emit BonusPaymentEvent(taskname, assignor, cost, _worker.getCreditScore(assignor),
        _applicant.getBalance(owner), _worker.getBalance(assignor));
    }
    
    //函数（数据查询相关）
    //--------1.获取申请者信息-----------
    function QueryApplicant() onlyApplicant public returns(address account, uint balance){
        return(msg.sender, _applicant.getBalance(msg.sender));
    }
    //--------2.获取工人信息-----------
    function QueryWorker() onlyWorker public returns(address account, uint balance, uint creditScore){
        return(msg.sender, _worker.getBalance(msg.sender), _worker.getCreditScore(msg.sender));
    }
    //--------3.获取发布的某个任务------------
    //input: 任务名
    //output: 任务名，任务发布者，任务持续时间，任务发布时间
    function QueryTask(string taskname) public view returns(string Taskname, address Owner, uint Duration,uint StartTime, string TaskBrief, uint ExpectedPrice){
        string memory tasknametemp;
        address owner;
        uint duration;      
        uint startTime;
        string memory taskBrief;
        uint expectedPrice;
        (tasknametemp, owner, duration, startTime, taskBrief, expectedPrice,,) = _task.getData(taskname);
        return(tasknametemp, owner, duration, startTime, taskBrief, expectedPrice);
    }
    //--------4.获取本申请者所有任务的任务名-------
    //结合getTask可获取我的所有任务信息
    function QueryAllMyTask() public returns(string[] tasknames){
        return _task.getAllMyTaskname();
    }
    
    //--------5.获取所有任务的任务名----------
    //结合getTask可得到所有任务信息，用于展示，方便工人选择
    function QueryAllTask() public view returns(string[] tasknames){
        return _task.getAllTaskname();
    }
    
    //---------6.获取本工人对某项任务的报价-----------
    //input: 任务名
    //output: 报价信息
    function QueryMyOffer(string taskname) onlyWorker public returns(uint ArrivalTime, uint DepartureTime, uint Cost, address Owner, string Taskname){
        uint arrivalTime;    
        uint departureTime;  
        uint cost;          
        address owner;       
        string memory tasknametemp;    
        (arrivalTime, departureTime, cost, owner, tasknametemp) = _offer.getData(taskname);
        return(arrivalTime, departureTime, cost, owner, tasknametemp);
    }
    //---------7.获取本工人所有的报价------------
    function QueryAllMyOffer() onlyWorker public returns(string[] tasknames){
        return _offer.getAllMyofferTotaskname();
    }
    
    //---------8.分配结果查询------------
    //input: 任务名
    //output: 任务名，被分配者，报价
    function QueryResult(string taskname) public returns(string Taskname, address Assignor, uint Cost){
        address assignor;
        uint cost;
        (,assignor,cost) = _result.getData(taskname);
        return(taskname, assignor, cost);
    }
}
