package main

import (
	"fmt"

	"os"

	"github.com/chainCompete/compete-service/blockchain"
)

func main() {
	// Definition of the Fabric SDK properties
	fSetup := blockchain.FabricSetup{
		// Network parameters
		OrdererID: "orderer.transaction",

		// Channel parameters
		ChannelID:     "datatransaction",
		ChannelConfig: os.Getenv("GOPATH") + "/src/github.com/chainCompete/compete-service/fixtures/artifacts/datatransaction.channel.tx",

		// Chaincode parameters
		ChainCodeID:     "transaction-service",
		ChaincodeGoPath: os.Getenv("GOPATH"),
		ChaincodePath:   "github.com/chainCompete/compete-service/chaincode/",
		OrgAdmin:        "Admin",
		OrgName:         "org1",
		ConfigFile:      "config.yaml",

		// User parameters
		UserName: "User1",
	}

	// Initialization of the Fabric SDK from the previously set properties
	err := fSetup.Initialize()
	if err != nil {
		fmt.Printf("Unable to initialize the Fabric SDK: %v\n", err)
		//return
	}
	// Close SDK
	defer fSetup.CloseSDK()

	// Install and instantiate the chaincode
	err = fSetup.InstallAndInstantiateCC()
	if err != nil {
		fmt.Printf("Unable to install and instantiate the chaincode: %v\n", err)
		//return
	}
	// app := &controllers.Application{
	// 	Fabric: &fSetup,
	// }
	// web.Serve(app)
	fSetup.RegisterUser("user1")
	fSetup.RegisterUser("user2")
	fSetup.RegisterUser("user3")
	fSetup.UserLogin("user1")
	fSetup.InvokeCreateUser("poster")
	fSetup.InvokePostTask("task1", "30", "2020-3-20", "test1", "20")
	fSetup.InvokePostTask("task2", "30", "2020-3-20", "test2", "20")
	fSetup.InvokePostTask("task3", "30", "2020-3-20", "test3", "20")
	fSetup.UserLogin("user2")
	fSetup.InvokeCreateUser("worker")
	fSetup.InvokePostOffer("2020", "2021", "10", "task1")
	fSetup.InvokePostOffer("2020", "2021", "20", "task2")
	fSetup.InvokePostOffer("2020", "2021", "30", "task3")
	fSetup.UserLogin("user3")
	fSetup.InvokeCreateUser("worker")
	fSetup.InvokePostOffer("2020", "2021", "30", "task1")
	fSetup.InvokePostOffer("2020", "2021", "20", "task2")
	fSetup.InvokePostOffer("2020", "2021", "10", "task3")
	fSetup.UserLogin("user1")
	fSetup.InvokeAssignTask("task1")
	fSetup.InvokeAssignTask("task2")
	fSetup.InvokeAssignTask("task3")
	fSetup.InvokeBonusPayment("task1","0")
	fSetup.InvokeBonusPayment("task2","0")
	fSetup.InvokeBonusPayment("task3","0")

	var choose int
	for {
		fmt.Println("1.Invoke, 2.Query, 3.Login, 4.Register, 5.Update")
		fmt.Scanf("%d", &choose)
		if choose == 1 {
			for {
				fmt.Println("Invoke")
				fmt.Scanf("%d", &choose)
				if choose == 1 {
					fmt.Println("InvokeCreateUser")
					var userRole string
					fmt.Scanln(&userRole)
					fmt.Println(fSetup.InvokeCreateUser(userRole))
				} else if choose == 2 {
					fmt.Println("InvokePostTask")
					var taskName, duration, startTime, taskBrief, exceptedPrice string
					fmt.Scanln(&taskName, &duration, &startTime, &taskBrief, &exceptedPrice)
					fmt.Println(fSetup.InvokePostTask(taskName, duration, startTime, taskBrief, exceptedPrice))
				} else if choose == 3 {
					fmt.Println("InvokePostOffer")
					var arrivalTime, departureTime, cost, taskName string
					fmt.Scanln(&arrivalTime, &departureTime, &cost, &taskName)
					fmt.Println(fSetup.InvokePostOffer(arrivalTime, departureTime, cost, taskName))
				} else if choose == 4 {
					fmt.Println("InvokeAssignTask")
					var taskName string
					fmt.Scanln(&taskName)
					fmt.Println(fSetup.InvokeAssignTask(taskName))
				} else if choose == 5 {
					fmt.Println("InvokeBonusPayment")
					var taskName,isSatisfied string
					fmt.Scanln(&taskName,&isSatisfied)
					fmt.Println(fSetup.InvokeBonusPayment(taskName,isSatisfied))
				} else {
					break
				}
			}
		} else if choose == 2 {
			for {
				fmt.Println("Query")
				fmt.Scanf("%d", &choose)
				if choose == 1 {
					fmt.Println("QueryUser")
					fmt.Println(fSetup.QueryUser())
				} else if choose == 2 {
					fmt.Println("QueryTask")
					var taskName string
					fmt.Scanln(&taskName)
					fmt.Println(fSetup.QueryTask(taskName))
				} else if choose == 3 {
					fmt.Println("QueryAllTask")
					fmt.Println(fSetup.QueryAllTask())
				} else if choose == 4 {
					fmt.Println("QueryUserTask")
					fmt.Println(fSetup.QueryUserTask())
				} else if choose == 5 {
					fmt.Println("QueryWorkerOffer")
					fmt.Println(fSetup.QueryWorkerOffer())
				} else if choose == 6 {
					fmt.Println("QueryWorkerTaskOffer")
					var taskName string
					fmt.Scanln(&taskName)
					fmt.Println(fSetup.QueryWorkerTaskOffer(taskName))
				} else if choose == 7 {
					fmt.Println("QueryAssignResult")
					var taskName string
					fmt.Scanln(&taskName)
					fmt.Println(fSetup.QueryAssignResult(taskName))
				} else {
					break
				}
			}
		} else if choose == 3 {
			var userName string
			fmt.Scanln(&userName)
			fSetup.UserLogin(userName)
		} else if choose == 4 {
			var userName string
			fmt.Scanln(&userName)
			err := fSetup.RegisterUser(userName)
			if err != nil {
				fmt.Println(err)
			}
		} else if choose == 5 {
			var version string
			fmt.Scanln(&version)
			fSetup.UpdateCC(version)
		}
	}

}

