package main

import (
	"encoding/json"
	"strconv"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

//CreateUser 用于存储用户信息
func (t *Chaincode) CreateUser(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	//args参数为用户角色
	if len(args) != 1 {
		return shim.Error("The number of args must be 1!")
	}
	//用户必须为poster或worker
	if args[0] != "poster" && args[0] != "worker" {
		return shim.Error("The role must be poster or worker!")
	}
	//获取调用者用户名
	uname, err := t.GetUserName(stub)
	if err != nil {
		return shim.Error(err.Error())
	}
	//判断用户是否已存在
	if t.IsUserExist(stub, uname) {
		return shim.Error("UserID already exist!")
	}
	user := User{UserID: uname, UserRole: args[0], UserAccount: 100}
	//将用户信息存储到区块中
	err = t.UserSetter(stub, user)
	if err != nil {
		return shim.Error(err.Error())
	}
	err = stub.SetEvent("eventInvokeCreateUser", []byte{}) //Notify listeners that an event "eventInvoke" have been executed
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(nil)
}

//PostTask 用于发布数据
func (t *Chaincode) PostTask(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	//args参数为任务名、持续时间、开始时间、任务概要、预期价格
	if len(args) != 5 {
		return shim.Error("The number of args must be 6!")
	}
	//获取调用者用户名
	uname, err := t.GetUserName(stub)
	if err != nil {
		return shim.Error(err.Error())
	}
	//判断用户是否为poster
	if !t.IsPoster(stub, uname) {
		return shim.Error("The invoker must be poster")
	}
	//判断任务是否存在
	if t.IsTaskExist(stub, args[0]) {
		return shim.Error("Task already exist!")
	}
	//存储任务信息到区块链
	task := Task{TaskName: args[0], StartTime: args[2], PosterID: uname, TaskBrief: args[3]}
	task.Duration, _ = strconv.Atoi(args[1])
	task.ExceptedPrice, _ = strconv.Atoi(args[4])

	err = t.TaskSetter(stub, task)
	if err != nil {
		return shim.Error(err.Error())
	}

	err = stub.SetEvent("eventInvokePostTask", []byte{}) //Notify listeners that an event "eventInvoke" have been executed
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(nil)
}

//PostOffer 用于工人发送请求
func (t *Chaincode) PostOffer(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	//args参数分别为到达时间、离开时间、报价、任务名
	if len(args) != 4 {
		return shim.Error("The number of args must be 4!")
	}
	//获取调用者用户名
	uname, err := t.GetUserName(stub)
	if err != nil {
		return shim.Error(err.Error())
	}
	//判断用户是否为worker
	if !t.IsWorker(stub, uname) {
		return shim.Error("The invoker must be worker")
	}
	//判断任务是否被分配
	if t.IsTaskEnd(stub, args[3]) {
		return shim.Error("Task has End!")
	}
	//判断是否已请求
	key, _ := stub.CreateCompositeKey("Offer", []string{args[3], uname})
	if t.IsOfferExist(stub, key) {
		return shim.Error("Offer has existed!")
	}
	//记录报价信息
	offer := Offer{ArrivalTime: args[0], DepartureTime: args[1], Worker: uname, TaskName: args[3]}
	offer.Cost, _ = strconv.Atoi(args[2])
	err = t.OfferSetter(stub, offer)
	if err != nil {
		return shim.Error(err.Error())
	}

	err = stub.SetEvent("eventInvokePostOffer", []byte{}) //Notify listeners that an event "eventInvoke" have been executed
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(nil)
}

//AssignTask 用于对任务进行分配
func (t *Chaincode) AssignTask(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	//args参数为任务名
	if len(args) != 1 {
		return shim.Error("The number of args must be 1!")
	}
	//获取调用者用户名
	uname, err := t.GetUserName(stub)
	if err != nil {
		return shim.Error(err.Error())
	}
	//判断用户是否为poster
	if !t.IsPoster(stub, uname) {
		return shim.Error("The invoker must be worker")
	}
	//查看任务内容
	task, err := t.TaskGetter(stub, args[0])
	if err != nil {
		return shim.Error(err.Error())
	}
	//任务是否已被分配
	if task.IsEnd {
		return shim.Error("The task has been assigned!")
	}
	//链码调用者是否为任务发布者
	if task.PosterID != uname {
		return shim.Error("The invoker must be task poster!")
	}
	//查看所有的请求
	rs, err := stub.GetStateByPartialCompositeKey("Offer", []string{args[0]})
	if err != nil {
		return shim.Error(err.Error())
	}
	data, err := GetListResult(rs)
	var list []Offer
	json.Unmarshal(data, &list)
	if len(list) == 0 {
		return shim.Error("No request!")
	}
	result := Result{TaskName: args[0]}
	//求出最低报价的用户以及对应的临界价值作为奖励值
	if len(list) == 1 {
		result.AssigneeID = list[0].Worker
		result.AssignCost = task.ExceptedPrice
	} else {
		mincost, bonus := 0, 0
		if list[0].Cost > list[1].Cost {
			mincost = 1
		} else {
			bonus = 1
		}
		for i := 2; i < len(list); i++ {
			if list[i].Cost < list[mincost].Cost {
				bonus = mincost
				mincost = i
			} else if list[i].Cost < list[bonus].Cost {
				bonus = i
			}
		}
		result.AssigneeID = list[mincost].Worker
		result.AssignCost = list[bonus].Cost
	}
	err = t.ResultSetter(stub, result)
	if err != nil {
		return shim.Error(err.Error())
	}
	//修改任务状态
	task.IsEnd = true
	err = t.TaskSetter(stub, task)
	if err != nil {
		return shim.Error(err.Error())
	}

	err = stub.SetEvent("eventInvokeAssignTask", []byte{}) //Notify listeners that an event "eventInvoke" have been executed
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(nil)
}

//BonusPayment 用于奖励支付
func (t *Chaincode) BonusPayment(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	//args参数为任务名
	if len(args) != 1 {
		return shim.Error("The number of args must be 1!")
	}
	//获取调用者用户名
	uname, err := t.GetUserName(stub)
	if err != nil {
		return shim.Error(err.Error())
	}
	//获取任务信息
	task, err := t.TaskGetter(stub, args[0])
	if err != nil {
		return shim.Error(err.Error())
	}
	//任务奖励是否已支付
	if task.IsPaid {
		return shim.Error("Task has been paid!")
	}
	//链码调用者是否为任务发布者
	if uname != task.PosterID {
		return shim.Error("The invoker must be the task poster!")
	}
	//获取分配结果
	result, err := t.ResultGetter(stub, args[0])
	if err != nil {
		return shim.Error(err.Error())
	}
	//获取任务发布者和任务分配者信息
	poster, err := t.UserGetter(stub, task.PosterID)
	if err != nil {
		return shim.Error(err.Error())
	}
	assignee, err := t.UserGetter(stub, result.AssigneeID)
	if err != nil {
		return shim.Error(err.Error())
	}
	//奖励支付
	poster.UserAccount -= result.AssignCost
	assignee.UserAccount += result.AssignCost
	//修改任务发布者和任务分配者状态
	err = t.UserSetter(stub, poster)
	if err != nil {
		return shim.Error(err.Error())
	}
	err = t.UserSetter(stub, assignee)
	if err != nil {
		return shim.Error(err.Error())
	}
	//修改任务状态
	task.IsPaid = true
	err = t.TaskSetter(stub, task)
	if err != nil {
		return shim.Error(err.Error())
	}

	err = stub.SetEvent("eventInvokeBonusPayment", []byte{}) //Notify listeners that an event "eventInvoke" have been executed
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(nil)
}
