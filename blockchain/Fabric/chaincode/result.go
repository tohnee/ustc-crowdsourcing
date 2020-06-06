package main

import (
	"encoding/json"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)

//分配结果信息
type Result struct {
	TaskName   string `json:TaskName`   //任务名
	AssigneeID string `json:AssigneeID` //被分配者
	AssignCost int    `json:AssignCost` //奖励支付
}

//ResultGetter 用于获取分配结果信息
func (t *Chaincode) ResultGetter(stub shim.ChaincodeStubInterface, taskName string) (Result, error) {
	key, _ := stub.CreateCompositeKey("TaskAssign", []string{taskName})
	resultJSONBytes, err := stub.GetState(key)
	result := Result{}
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(resultJSONBytes, &result)
	if err != nil {
		return result, err
	}
	return result, err
}

//ResultSetter 用于存储分配结果信息
func (t *Chaincode) ResultSetter(stub shim.ChaincodeStubInterface, result Result) error {
	resultJSONBytes, err := json.Marshal(result)
	if err != nil {
		return err
	}
	key, _ := stub.CreateCompositeKey("TaskAssign", []string{result.TaskName})
	err = stub.PutState(key, resultJSONBytes)
	return err
}
