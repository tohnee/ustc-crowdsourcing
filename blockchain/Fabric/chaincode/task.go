package main

import (
	"encoding/json"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)

//任务信息
type Task struct {
	TaskName      string `json:TaskName`      //任务名
	Duration      int    `json:Duration`      //任务持续时间
	StartTime     string `json:StartTime`     //任务开始时间
	PosterID      string `json:Owner`         //任务发布者ID
	TaskBrief     string `json:TaskBrief`     //任务简介
	IsEnd         bool   `json:IsEnd`         //任务是否结束
	IsPaid        bool   `json:IsPaid`        //是否给予奖励
	ExceptedPrice int    `json:ExceptedPrice` //用户自定任务价格
}

//IsTaskExist 判断任务是否存在
func (t *Chaincode) IsTaskExist(stub shim.ChaincodeStubInterface, taskName string) bool {
	taskJSONBytes, err := stub.GetState(taskName)
	return err != nil || taskJSONBytes != nil
}

//IsTaskEnd 判断任务是否结束
func (t *Chaincode) IsTaskEnd(stub shim.ChaincodeStubInterface, taskID string) bool {
	taskJSONBytes, err := stub.GetState(taskID)
	if err != nil {
		return true
	}
	task := Task{}
	err = json.Unmarshal(taskJSONBytes, &task)
	if err != nil {
		return true
	}
	return task.IsEnd
}

//IsTaskPaid 判断任务是否给予奖励
func (t *Chaincode) IsTaskPaid(stub shim.ChaincodeStubInterface, taskID string) bool {
	taskJSONBytes, err := stub.GetState(taskID)
	if err != nil {
		return true
	}
	task := Task{}
	err = json.Unmarshal(taskJSONBytes, &task)
	if err != nil {
		return true
	}
	return task.IsPaid
}

//TaskGetter 通过任务名获取任务结构体
func (t *Chaincode) TaskGetter(stub shim.ChaincodeStubInterface, taskName string) (Task, error) {
	taskJSONBytes, err := stub.GetState(taskName)
	task := Task{}
	if err != nil {
		return task, err
	}
	err = json.Unmarshal(taskJSONBytes, &task)
	if err != nil {
		return task, err
	}
	return task, nil
}

//TaskSetter 将任务信息存入区块链中
func (t *Chaincode) TaskSetter(stub shim.ChaincodeStubInterface, task Task) error {
	taskJSONBytes, err := json.Marshal(task)
	if err != nil {
		return err
	}
	key, _ := stub.CreateCompositeKey("Task", []string{task.PosterID, task.TaskName})
	err = stub.PutState(key, taskJSONBytes)
	if err != nil {
		return err
	}
	err = stub.PutState(task.TaskName, taskJSONBytes)
	return err
}
