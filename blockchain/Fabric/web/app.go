package web

import (
	"fmt"
	"net/http"

	"github.com/chainCompete/compete-service/web/controllers"
)


// Serve starts
func Serve(app *controllers.Application) {
	mux := http.NewServeMux()

	mux.HandleFunc("/register", app.PostRegisterRequest)
	mux.HandleFunc("/login", app.PostLoginRequest)
	mux.HandleFunc("/posttask", app.PostTask)
	mux.HandleFunc("/postoffer", app.PostOffer)
	mux.HandleFunc("/assigntask", app.AssignTask)
	mux.HandleFunc("/bonuspayment", app.BonusPayment)

	mux.HandleFunc("/queryuser", app.QueryUser)
	mux.HandleFunc("/querytask", app.QueryTask)
	mux.HandleFunc("/queryalltask", app.QueryAllTask)
	mux.HandleFunc("/queryusertask", app.QueryUserTask)
	mux.HandleFunc("/queryworkeroffer", app.QueryWorkerOffer)
	mux.HandleFunc("/queryworkertaskoffer", app.QueryWorkerTaskOffer)
	mux.HandleFunc("/queryassignresult", app.QueryAssignResult)

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/error", http.StatusTemporaryRedirect)
	})

	fmt.Println("Listening (http://localhost:3000/) ...")
	http.ListenAndServe(":3000", mux)
}

