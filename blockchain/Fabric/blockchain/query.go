package blockchain

import (
	"fmt"

	"github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
)

//QueryUser 查询用户信息
func (setup *FabricSetup) QueryUser() (string, error) {

	response, err := setup.client.Query(channel.Request{ChaincodeID: setup.ChainCodeID, Fcn: "QueryUser", Args: [][]byte{}})
	if err != nil {
		return "", fmt.Errorf("failed to query: %v", err)
	}
	return string(response.Payload), nil
}

//QueryTask 查询任务信息
func (setup *FabricSetup) QueryTask(taskID string) (string, error) {

	response, err := setup.client.Query(channel.Request{ChaincodeID: setup.ChainCodeID, Fcn: "QueryTask", Args: [][]byte{[]byte(taskID)}})
	if err != nil {
		return "", fmt.Errorf("failed to query: %v", err)
	}
	return string(response.Payload), nil
}

//QueryAllTask 查看所有任务信息
func (setup *FabricSetup) QueryAllTask() (string, error) {

	response, err := setup.client.Query(channel.Request{ChaincodeID: setup.ChainCodeID, Fcn: "QueryAllTask", Args: [][]byte{}})
	if err != nil {
		return "", fmt.Errorf("failed to query: %v", err)
	}
	return string(response.Payload), nil
}

//QueryUserTask 查看某个用户发布的任务信息
func (setup *FabricSetup) QueryUserTask() (string, error) {

	response, err := setup.client.Query(channel.Request{ChaincodeID: setup.ChainCodeID, Fcn: "QueryUserTask", Args: [][]byte{}})
	if err != nil {
		return "", fmt.Errorf("failed to query: %v", err)
	}

	return string(response.Payload), nil
}

//QueryWorkerOffer 查看工人提出的所有请求
func (setup *FabricSetup) QueryWorkerOffer() (string, error) {

	response, err := setup.client.Query(channel.Request{ChaincodeID: setup.ChainCodeID, Fcn: "QueryWorkerOffer", Args: [][]byte{}})
	if err != nil {
		return "", fmt.Errorf("failed to query: %v", err)
	}

	return string(response.Payload), nil
}

//QueryUserData 查询工人对具体某个任务
func (setup *FabricSetup) QueryWorkerTaskOffer(taskName string) (string, error) {

	response, err := setup.client.Query(channel.Request{ChaincodeID: setup.ChainCodeID, Fcn: "QueryWorkerTaskOffer", Args: [][]byte{[]byte(taskName)}})
	if err != nil {
		return "", fmt.Errorf("failed to query: %v", err)
	}

	return string(response.Payload), nil
}

//QueryAssignResult 查询某个任务的分配结果
func (setup *FabricSetup) QueryAssignResult(taskName string) (string, error) {

	response, err := setup.client.Query(channel.Request{ChaincodeID: setup.ChainCodeID, Fcn: "QueryAssignResult", Args: [][]byte{[]byte(taskName)}})
	if err != nil {
		return "", fmt.Errorf("failed to query: %v", err)
	}

	return string(response.Payload), nil
}
