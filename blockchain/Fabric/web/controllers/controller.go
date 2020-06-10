package controllers

import (
	"io"
	"net/http"
	"github.com/chainCompete/compete-service/blockchain"
)

//Application represents that
type Application struct {
	Fabric *blockchain.FabricSetup
}

func (app *Application) PostRegisterRequest(w http.ResponseWriter, r *http.Request) {
	userName := r.PostFormValue("username")
	userRole := r.PostFormValue("userrole")
	if userName == "" || userRole == "" {
		if userName == ""{
			io.WriteString(w, "username is empty!")
		}
		if userRole == ""{
			io.WriteString(w, "userrole is empty!")		
		}
	}else {
		success := true
		err := app.Fabric.RegisterUser(userName)
		if err != nil {
			success = false
		}
		err = app.Fabric.UserLogin(userName)
		if err != nil {
			success = false
		}
		_, err = app.Fabric.InvokeCreateUser(userRole)
		if err != nil {
			success = false
		}
		if success {
			io.WriteString(w, "注册成功")
		} else {
			io.WriteString(w, "注册失败")
		}
	}
}
func (app *Application) PostLoginRequest(w http.ResponseWriter, r *http.Request) {
	userName := r.PostFormValue("username")
	err := app.Fabric.UserLogin(userName)
	if err == nil {
		io.WriteString(w, "登陆成功")
	} else {
		io.WriteString(w, "登陆失败")
	}

}
func (app *Application) PostTask(w http.ResponseWriter, r *http.Request) {
	taskName := r.PostFormValue("taskname")
	duration := r.PostFormValue("duration")
	startTime := r.PostFormValue("starttime")
	taskBrif := r.PostFormValue("taskbrif")
	exceptedPrice := r.PostFormValue("exceptedprice")
	_, err := app.Fabric.InvokePostTask(taskName, duration, startTime, taskBrif, exceptedPrice)
	if err == nil {
		io.WriteString(w, "发布成功")
	} else {
		io.WriteString(w, "发布失败")
	}

}
func (app *Application) PostOffer(w http.ResponseWriter, r *http.Request) {

	arrivalTime := r.PostFormValue("arrivaltime")
	departureTime := r.PostFormValue("departuretime")
	cost := r.PostFormValue("cost")
	taskName := r.PostFormValue("taskname")
	_, err := app.Fabric.InvokePostOffer(arrivalTime, departureTime, cost, taskName)
	if err == nil {
		io.WriteString(w, "请求成功")
	} else {
		io.WriteString(w, "请求失败")
	}

}
func (app *Application) AssignTask(w http.ResponseWriter, r *http.Request) {
	taskName := r.PostFormValue("taskname")
	_, err := app.Fabric.InvokeAssignTask(taskName)
	if err == nil {
		io.WriteString(w, "分配成功")
	} else {
		io.WriteString(w, "分配失败")
	}

}
func (app *Application) BonusPayment(w http.ResponseWriter, r *http.Request) {
	taskName := r.PostFormValue("taskname")
	isSatisfied := r.PostFormValue("issatisfied")
	_, err := app.Fabric.InvokeBonusPayment(taskName, isSatisfied)
	if err == nil {
		io.WriteString(w, "支付成功")
	} else {
		io.WriteString(w, "支付失败")
	}

}
func (app *Application) QueryUser(w http.ResponseWriter, r *http.Request) {
	data, _ := app.Fabric.QueryUser()
	io.WriteString(w, data)
}
func (app *Application) QueryTask(w http.ResponseWriter, r *http.Request) {
	taskName := r.PostFormValue("taskname")
	data, _ := app.Fabric.QueryTask(taskName)
	io.WriteString(w, data)
}
func (app *Application) QueryAllTask(w http.ResponseWriter, r *http.Request) {
	data, _ := app.Fabric.QueryAllTask()
	io.WriteString(w, data)
}
func (app *Application) QueryUserTask(w http.ResponseWriter, r *http.Request) {
	data, _ := app.Fabric.QueryUserTask()
	io.WriteString(w, data)
}
func (app *Application) QueryWorkerOffer(w http.ResponseWriter, r *http.Request) {
	data, _ := app.Fabric.QueryWorkerOffer()
	io.WriteString(w, data)
}

func (app *Application) QueryWorkerTaskOffer(w http.ResponseWriter, r *http.Request) {
	taskName := r.PostFormValue("taskname")
	data, _ := app.Fabric.QueryWorkerTaskOffer(taskName)
	io.WriteString(w, data)

}
func (app *Application) QueryAssignResult(w http.ResponseWriter, r *http.Request) {
	taskName := r.PostFormValue("taskname")
	data, _ := app.Fabric.QueryAssignResult(taskName)
	io.WriteString(w, data)

}

