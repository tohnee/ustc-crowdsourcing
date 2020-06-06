package main

import (
	"encoding/json"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)

//工人报价信息
type Offer struct {
	ArrivalTime   string `json:ArrivalTime`   //到达时间
	DepartureTime string `json:DepartureTime` //离开时间
	Cost          int    `json:Cost`          //报价
	Worker        string `json:Worker`        //工人名
	TaskName      string `json:TaskName`      //任务名
}

//IsOfferExist 判断请求是否存在
func (t *Chaincode) IsOfferExist(stub shim.ChaincodeStubInterface, key string) bool {
	offerJSONBytes, err := stub.GetState(key)
	return err != nil || offerJSONBytes != nil
}

//OfferSetter 将工人请求存入区块链
func (t *Chaincode) OfferSetter(stub shim.ChaincodeStubInterface, offer Offer) error {
	key, _ := stub.CreateCompositeKey("Offer", []string{offer.TaskName, offer.Worker})
	offerJSONBytes, err := json.Marshal(offer)
	if err != nil {
		return err
	}
	err = stub.PutState(key, offerJSONBytes)
	if err != nil {
		return err
	}
	key, _ = stub.CreateCompositeKey("OfferWorker", []string{offer.Worker, offer.TaskName})
	err = stub.PutState(key, offerJSONBytes)
	return err
}
