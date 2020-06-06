package main

import (
	"bytes"
	"crypto/x509"
	"encoding/pem"
	"errors"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

//GetListResult 将含有多个的查询结果转换为JSON格式字符串
func GetListResult(resultsIterator shim.StateQueryIteratorInterface) ([]byte, error) {

	defer resultsIterator.Close()
	//将查询结果以Json字符串的格式存储到buffer中
	var buffer bytes.Buffer
	buffer.WriteString("[")

	bArrayMemberAlreadyWritten := false
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}
		if bArrayMemberAlreadyWritten == true {
			buffer.WriteString(",")
		}
		buffer.WriteString(string(queryResponse.Value))
		bArrayMemberAlreadyWritten = true
	}
	buffer.WriteString("]")
	return buffer.Bytes(), nil
}

//GetUserName 用于查看用户名
func (t *Chaincode) GetUserName(stub shim.ChaincodeStubInterface) (string, error) {
	creatorByte, _ := stub.GetCreator()
	certStart := bytes.IndexAny(creatorByte, "-----BEGIN")
	if certStart == -1 {
		return "", errors.New("No certificate found")
	}
	certText := creatorByte[certStart:]
	bl, _ := pem.Decode(certText)
	if bl == nil {
		return "", errors.New("Could not decode the PEM structure")
	}

	cert, err := x509.ParseCertificate(bl.Bytes)
	if err != nil {
		return "", errors.New("ParseCertificate failed")
	}
	return cert.Subject.CommonName, nil
}

//QueryUser 用于查询个人信息
func (t *Chaincode) QueryUser(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	//获取调用者用户名
	uname, err := t.GetUserName(stub)
	if err != nil {
		return shim.Error(err.Error())
	}
	userJSONBytes, err := stub.GetState(uname)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(userJSONBytes)
}

//QueryTask 可以通过任务名查询任务信息
func (t *Chaincode) QueryTask(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 1 {
		return shim.Error("The number of args must be 1!")
	}
	taskJSONBytes, err := stub.GetState(args[0])
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(taskJSONBytes)
}

//QueryAllTask 可以通过任务名查询任务信息
func (t *Chaincode) QueryAllTask(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	rs, err := stub.GetStateByPartialCompositeKey("Task", []string{})
	if err != nil {
		return shim.Error(err.Error())
	}
	data, err := GetListResult(rs)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(data)
}

//QueryUserTask 用于查询个人发布的所有任务
func (t *Chaincode) QueryUserTask(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	//获取调用者用户名
	uname, err := t.GetUserName(stub)
	if err != nil {
		return shim.Error(err.Error())
	}
	//获取调用者发布的所有任务
	rs, err := stub.GetStateByPartialCompositeKey("Task", []string{uname})
	if err != nil {
		return shim.Error(err.Error())
	}
	data, err := GetListResult(rs)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(data)
}

//QueryWorkerOffer 用于查询工人的所有请求
func (t *Chaincode) QueryWorkerOffer(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	//获取调用者用户名
	uname, err := t.GetUserName(stub)
	if err != nil {
		return shim.Error(err.Error())
	}
	//获取本工人的所有请求信息
	rs, err := stub.GetStateByPartialCompositeKey("OfferWorker", []string{uname})
	if err != nil {
		return shim.Error(err.Error())
	}
	data, err := GetListResult(rs)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(data)
}

//QueryWorkerTaskOffer 可以本工人对某个任务的请求
func (t *Chaincode) QueryWorkerTaskOffer(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 1 {
		return shim.Error("The number of args must be 1!")
	}
	//获取调用者的用户名
	uname, err := t.GetUserName(stub)
	if err != nil {
		return shim.Error(err.Error())
	}
	//查询个人对某个任务的请求
	key, _ := stub.CreateCompositeKey("OfferWorker", []string{uname, args[0]})
	offerJSONBytes, err := stub.GetState(key)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(offerJSONBytes)
}

//QueryAssignResult 查询任务分配结果
func (t *Chaincode) QueryAssignResult(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 1 {
		return shim.Error("The number of args must be 1!")
	}
	//根据任务名查询任务的分配结果
	key, _ := stub.CreateCompositeKey("TaskAssign", args)
	resultJSONBytes, err := stub.GetState(key)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(resultJSONBytes)
}
