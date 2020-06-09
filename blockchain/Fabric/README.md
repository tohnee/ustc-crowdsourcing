#### 角色分配

***

##### Admin 

* 发布合约、更新合约

##### Poster

* 发布任务、分配任务、奖励支付

##### Worker

* 参与报价

#### 层次结构

***

![](https://github.com/tohnee/ustc-crowdsourcing/raw/master/blockchain/Fabric/images/1.png)

数据层表示存储在Fabric中的相关信息，链码层提供对区块链中数据的查询和写入，蓝色代表写交易，红色代表读交易，中间层负责链码函数调用，用户准入的控制和权限控制，以及链码的安装和更新。

* chaincode/chaincodemain.go 主程序入口，解析用户输入调用不同链码 
* chaincode/chaincodequery.go 与查询相关的所有链码函数
* chaincode/chaincodeinvoke.go 与数据写入相关的所有链码函数
* chaincode/task.go、chaincode/user.go、chaincode/offer.go、chaincode/result.go 定义任务、用户、工人请求、分配结果的结构体，并提供存取的函数

#### 链码接口

***

##### 链码入口

```go
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
```

##### 链码函数（数据写入）

* func (t *Chaincode) CreateUser(stub shim.ChaincodeStubInterface, args []string)  

  CreateUser用于存储用户信息，args参数为用户角色信息（必须为poster或worker），当用户信息不存在时，将用户信息存储到区块链中，否则返回错误。

* func (t *Chaincode) PostTask(stub shim.ChaincodeStubInterface, args []string)  

  PostTask用于发布任务，链码调用者必须为poster，args参数为任务名、持续时间、开始时间、任务概要、预期价格，当任务未被发布时，将任务信息存储到区块链中，否则返回错误。

* func (t *Chaincode) PostOffer(stub shim.ChaincodeStubInterface, args []string)  

  PostOffer用于工人发送请求，链码调用者必须为worker，args参数分别为到达时间、离开时间、报价、任务名，当任务未被分配且工人未请求该任务时，将请求信息存储到区块链中，否则返回错误。

* func (t *Chaincode) AssignTask(stub shim.ChaincodeStubInterface, args []string) 

  AssignTask用于对任务进行分配，链码调用者必须为任务的发布者，args参数为任务名，当任务未被分配且存在工人请求时，系统会将任务分配给报价最低的工人，计算临界价值作为工人的奖励，并把该分配结果存储区块链中，否则返回错误。

* func (t *Chaincode) BonusPayment(stub shim.ChaincodeStubInterface, args []string) 

  BonusPayment用于奖励支付，链码调用者必须为任务的发布者，args参数为任务名和满意度，当任务对应的分配结果存在时，系统会进行积分的转移，当发布者不满足工人的工作时，工人的信誉度会下降，得到的报酬也会降低，否则返回错误。

##### 链码函数（数据查询）

* func (t *Chaincode) QueryUser(stub shim.ChaincodeStubInterface, args []string) 

  QueryUser用于查询链码调用者的个人信息，不需要参数。

* func (t *Chaincode) QueryTask(stub shim.ChaincodeStubInterface, args []string) 

  QueryTask通过任务名查询系统中的任务信息，args参数为任务名。

* func (t *Chaincode) QueryAllTask(stub shim.ChaincodeStubInterface, args []string)  

  QueryAllTask用于查询系统中发布的所有任务信息，不需要参数。

* func (t *Chaincode) QueryUserTask(stub shim.ChaincodeStubInterface, args []string)  

  QueryUserTask用于查询链码调用者发布的所有任务，不需要参数。

* func (t *Chaincode) QueryWorkerOffer(stub shim.ChaincodeStubInterface, args []string) 

  QueryWorkerOffer用于查询链码调用者的所有任务请求信息，不需要参数。

* func (t *Chaincode) QueryWorkerTaskOffer(stub shim.ChaincodeStubInterface, args []string)  

  QueryWorkerTaskOffer用于链码调用者对某个任务的请求信息，args参数为任务名。

* func (t *Chaincode) QueryAssignResult(stub shim.ChaincodeStubInterface, args []string)  

  QueryAssignResult用于查询任务分配结果信息，args参数为任务名。

#### 中间层接口

***

##### 用户准入

* func (setup *FabricSetup) RegisterUser(userName string) error 
* func (setup *FabricSetup) UserLogin(userName string) error

| 接口         | 参数             | 返回值 | 描述     |
| ------------ | ---------------- | ------ | -------- |
| RegisterUser | userName(string) | error  | 用户注册 |
| UserLogin    | userName(string) | error  | 用户登陆 |

##### 链码安装/更新

* func (setup *FabricSetup) InstallAndInstantiateCC() error 
* func (setup *FabricSetup) UpdateCC(version string) error

| 接口                    | 参数            | 返回值 | 描述             |
| ----------------------- | --------------- | ------ | ---------------- |
| InstallAndInstantiateCC | 空              | error  | 链码安装及初始化 |
| UpdateCC                | version(string) | error  | 链码更新         |

##### 链码调用（数据写入相关）

* func (setup *FabricSetup) InvokeCreateUser(userRole string) (string, error) 

* func (setup *FabricSetup) InvokePostTask(taskName, duration, startTime,  taskBrief, exceptedPrice string) (string, error) 

* func (setup *FabricSetup) InvokePostOffer(arrivalTime, departureTime, cost,  taskName string) (string, error) 

* func (setup *FabricSetup) InvokeAssignTask(taskName string) (string, error) 

* func (setup *FabricSetup) InvokeBonusPayment(taskName string) (string, error) 

| 接口               | 参数                                                         | 返回值        | 描述             |
| ------------------ | ------------------------------------------------------------ | ------------- | ---------------- |
| InvokeCreateUser   | userRole(string)                                             | string, error | 存储用户相关信息 |
| InvokePostTask     | taskName, duration, startTime, taskBrief, exceptedPrice(string) | string, error | 发布任务         |
| InvokePostOffer    | arrivalTime, departureTime, cost, taskName(string)           | string, error | 工人报价请求     |
| InvokeAssignTask   | taskName(string)                                             | string, error | 任务分配         |
| InvokeBonusPayment | taskName(string)                                             | string, error | 奖励支付         |

##### 链码调用（查询相关）

* func (setup *FabricSetup) QueryUser() (string, error) 

* func (setup *FabricSetup) QueryTask(taskID string) (string, error) 

* func (setup *FabricSetup) QueryAllTask() (string, error) 

* func (setup *FabricSetup) QueryUserTask() (string, error) 

* func (setup *FabricSetup) QueryWorkerOffer() (string, error)

* func (setup *FabricSetup) QueryWorkerTaskOffer(taskName string) (string, error) 

* func (setup *FabricSetup) QueryAssignResult(taskName string) (string, error) 

| 接口                 | 参数             | 返回值       | 描述                         |
| -------------------- | ---------------- | ------------ | ---------------------------- |
| GetUser              | 无               | string,error | 返回调用合约的用户名         |
| QueryUser            | 无               | string,error | 查询用户相关信息             |
| QueryTask            | taskID(string)   | string,error | 查询任务相关信息             |
| QueryAllTask         | 无               | string,error | 查询系统中发布的所有任务     |
| QueryUserTask        | 无               | string,error | 查询某个用户发布的所有任务   |
| QueryWorkerOffer     | 无               | string,error | 查询某个工人的所有请求       |
| QueryWorkerTaskOffer | 无               | string,error | 查询某个工人对某个任务的请求 |
| QueryAssignResult    | taskName(string) | string,error | 查询某个任务的分配结果       |

#### 链码测试（通过中间层Fabric Go SDK调用）

在测试样例中，注册了三个用户，user1，user2，user3 ，user1为任务发布者，user2，user3为工人。user1发布了三个任务分别是task1，task2，task3，user2和user3分别对这三个任务发起了请求，user1在接受到请求后对任务进行了分配和奖励支付。测试程序如下：

```go
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
	fSetup.InvokePostOffer("2020", "2021", "30", "task1")
	fSetup.InvokePostOffer("2020", "2021", "20", "task2")
	fSetup.InvokePostOffer("2020", "2021", "10", "task3")
	fSetup.UserLogin("user3")
	fSetup.InvokeCreateUser("worker")
	fSetup.InvokePostOffer("2020", "2021", "10", "task1")
	fSetup.InvokePostOffer("2020", "2021", "20", "task2")
	fSetup.InvokePostOffer("2020", "2021", "30", "task3")
	fSetup.UserLogin("user1")
	fSetup.InvokeAssignTask("task1")
	fSetup.InvokeAssignTask("task2")
	fSetup.InvokeAssignTask("task3")
	fSetup.InvokeBonusPayment("task1")
	fSetup.InvokeBonusPayment("task2")
	fSetup.InvokeBonusPayment("task3")
```

链码正确执行，程序执行结果如下：

![image-20200607083704425](https://github.com/tohnee/ustc-crowdsourcing/raw/master/blockchain/Fabric/images/2.png)

对区块链中的数据进行查询，结果如下：

* 查询用户信息（以user1为例）

![image-20200607084355175](https://github.com/tohnee/ustc-crowdsourcing/raw/master/blockchain/Fabric/images/3.png)

* 查询任务信息（以task1为例）

![image-20200607084822395](https://github.com/tohnee/ustc-crowdsourcing/raw/master/blockchain/Fabric/images/4.png)

* 查询系统中的所有任务

![image-20200607084757513](https://github.com/tohnee/ustc-crowdsourcing/raw/master/blockchain/Fabric/images/5.png)

* 查询个人发布的所有任务（以user1为例）

![image-20200607084907263](https://github.com/tohnee/ustc-crowdsourcing/raw/master/blockchain/Fabric/images/6.png)

* 查询个人发布的请求（以user2为例）

![image-20200607085240220](https://github.com/tohnee/ustc-crowdsourcing/raw/master/blockchain/Fabric/images/7.png)

* 查询个人对某个任务的请求（以user2为例）

![image-20200607085353045](https://github.com/tohnee/ustc-crowdsourcing/raw/master/blockchain/Fabric/images/8.png)

* 查询某个任务的分配结果（以task1，task2，task3为例）

![image-20200607085501738](https://github.com/tohnee/ustc-crowdsourcing/raw/master/blockchain/Fabric/images/9.png)
