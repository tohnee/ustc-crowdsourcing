package blockchain

import (
	"fmt"
	"time"

	"github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
)

// InvokeCreateUser 调用链码函数CreateUser
func (setup *FabricSetup) InvokeCreateUser(userRole string) (string, error) {

	eventID := "eventInvokeCreateUser"

	reg, notifier, err := setup.event.RegisterChaincodeEvent(setup.ChainCodeID, eventID)
	if err != nil {
		return "", err
	}
	defer setup.event.Unregister(reg)

	response, err := setup.client.Execute(channel.Request{ChaincodeID: setup.ChainCodeID, Fcn: "CreateUser", Args: [][]byte{[]byte(userRole)}})

	if err != nil {
		return "", fmt.Errorf("failed to move funds: %v", err)
	}
	select {
	case ccEvent := <-notifier:
		fmt.Printf("Received CC event: %v\n", ccEvent)
	case <-time.After(time.Second * 20):
		return "", fmt.Errorf("did NOT receive CC event for eventId(%s)", eventID)
	}

	return string(response.Payload), nil
}

// InvokePostTask
func (setup *FabricSetup) InvokePostTask(taskName, duration, startTime, taskBrief, exceptedPrice string) (string, error) {

	eventID := "eventInvokePostTask"

	reg, notifier, err := setup.event.RegisterChaincodeEvent(setup.ChainCodeID, eventID)
	if err != nil {
		return "", err
	}
	defer setup.event.Unregister(reg)

	response, err := setup.client.Execute(channel.Request{ChaincodeID: setup.ChainCodeID, Fcn: "PostTask", Args: [][]byte{[]byte(taskName), []byte(duration), []byte(startTime), []byte(taskBrief), []byte(exceptedPrice)}})

	if err != nil {
		return "", fmt.Errorf("failed to move funds: %v", err)
	}
	select {
	case ccEvent := <-notifier:
		fmt.Printf("Received CC event: %v\n", ccEvent)
	case <-time.After(time.Second * 20):
		return "", fmt.Errorf("did NOT receive CC event for eventId(%s)", eventID)
	}

	return string(response.Payload), nil
}

// InvokePostOffer
func (setup *FabricSetup) InvokePostOffer(arrivalTime, departureTime, cost, taskName string) (string, error) {

	eventID := "eventInvokePostOffer"

	reg, notifier, err := setup.event.RegisterChaincodeEvent(setup.ChainCodeID, eventID)
	if err != nil {
		return "", err
	}
	defer setup.event.Unregister(reg)

	response, err := setup.client.Execute(channel.Request{ChaincodeID: setup.ChainCodeID, Fcn: "PostOffer", Args: [][]byte{[]byte(arrivalTime), []byte(departureTime), []byte(cost), []byte(taskName)}})

	if err != nil {
		return "", fmt.Errorf("failed to move funds: %v", err)
	}
	select {
	case ccEvent := <-notifier:
		fmt.Printf("Received CC event: %v\n", ccEvent)
	case <-time.After(time.Second * 20):
		return "", fmt.Errorf("did NOT receive CC event for eventId(%s)", eventID)
	}

	return string(response.Payload), nil
}

// InvokeAssignTask
func (setup *FabricSetup) InvokeAssignTask(taskName string) (string, error) {

	eventID := "eventInvokeAssignTask"

	reg, notifier, err := setup.event.RegisterChaincodeEvent(setup.ChainCodeID, eventID)
	if err != nil {
		return "", err
	}
	defer setup.event.Unregister(reg)

	response, err := setup.client.Execute(channel.Request{ChaincodeID: setup.ChainCodeID, Fcn: "AssignTask", Args: [][]byte{[]byte(taskName)}})

	if err != nil {
		return "", fmt.Errorf("failed to move funds: %v", err)
	}
	select {
	case ccEvent := <-notifier:
		fmt.Printf("Received CC event: %v\n", ccEvent)
	case <-time.After(time.Second * 20):
		return "", fmt.Errorf("did NOT receive CC event for eventId(%s)", eventID)
	}

	return string(response.Payload), nil
}

// InvokeBonusPayment
func (setup *FabricSetup) InvokeBonusPayment(taskName, isSatisfied string) (string, error) {

	eventID := "eventInvokeBonusPayment"

	reg, notifier, err := setup.event.RegisterChaincodeEvent(setup.ChainCodeID, eventID)
	if err != nil {
		return "", err
	}
	defer setup.event.Unregister(reg)

	response, err := setup.client.Execute(channel.Request{ChaincodeID: setup.ChainCodeID, Fcn: "BonusPayment", Args: [][]byte{[]byte(taskName), []byte(isSatisfied)}})

	if err != nil {
		return "", fmt.Errorf("failed to move funds: %v", err)
	}
	select {
	case ccEvent := <-notifier:
		fmt.Printf("Received CC event: %v\n", ccEvent)
	case <-time.After(time.Second * 20):
		return "", fmt.Errorf("did NOT receive CC event for eventId(%s)", eventID)
	}

	return string(response.Payload), nil
}
