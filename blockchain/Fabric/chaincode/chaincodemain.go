package main

import (
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

type Chaincode struct {
}

//初始化链码
func (t *Chaincode) Init(stub shim.ChaincodeStubInterface) pb.Response {
	return shim.Success(nil)
}

//链码调用
func (t *Chaincode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	function, args := stub.GetFunctionAndParameters()
	if function == "CreateUser" { //创建用户
		return t.CreateUser(stub, args)
	} else if function == "PostTask" { //发布任务
		return t.PostTask(stub, args)
	} else if function == "PostOffer" { //工人请求
		return t.PostOffer(stub, args)
	} else if function == "AssignTask" { //分配任务
		return t.AssignTask(stub, args)
	} else if function == "BonusPayment" { //奖励支付
		return t.BonusPayment(stub, args)
	} else if function == "QueryUser" { //查询用户
		return t.QueryUser(stub, args)
	} else if function == "QueryTask" { //查询任务
		return t.QueryTask(stub, args)
	} else if function == "QueryAllTask" { //查询所有任务
		return t.QueryAllTask(stub, args)
	} else if function == "QueryUserTask" { //查询用户任务
		return t.QueryUserTask(stub, args)
	} else if function == "QueryWorkerOffer" { //查询本工人的所有请求
		return t.QueryWorkerOffer(stub, args)
	} else if function == "QueryWorkerTaskOffer" { //查询本工人对某个任务的请求
		return t.QueryWorkerTaskOffer(stub, args)
	} else if function == "QueryAssignResult" { //查询任务分配结果
		return t.QueryAssignResult(stub, args)
	}
	return shim.Error("Received unknown function invocation")
}

func main() {
	err := shim.Start(new(Chaincode))
	if err != nil {
		fmt.Printf("Error starting Trade chaincode: %s ", err)
	}
}
