package routes

import (
	"swtest2/internal/transport/handlers"

	"github.com/gorilla/mux"
)

var ColleagueManagerRoutes = func(router *mux.Router) {
	router.HandleFunc("/colleaguemanager/addcolleague", handlers.AddColleague).Methods("POST")
	router.HandleFunc("/colleaguemanager/blockcolleague/{id}", handlers.BlockColleague).Methods("GET")
	router.HandleFunc("/colleaguemanager/addclient", handlers.AddClient).Methods("POST")
	router.HandleFunc("/colleaguemanager/showclientinfo/{id}", handlers.ShowClientInfo).Methods("GET")
	router.HandleFunc("/colleaguemanager/changeclient/{id}", handlers.ChangeClient).Methods("PUT")
	router.HandleFunc("/colleaguemanager/showclientsapplicationsinfo", handlers.ShowClientApplicationsInfo).Methods("GET")
	router.HandleFunc("/colleaguemanager/changeapplicationstatus/{id}", handlers.ChangeApplicationStatus).Methods("PUT")
	router.HandleFunc("/colleaguemanager/addcompany", handlers.AddCompany).Methods("POST")
	router.HandleFunc("/colleaguemanager/addclientintocompany", handlers.AddClientIntoCompany).Methods("PUT")
	router.HandleFunc("/colleaguemanager/changecompany/{id}", handlers.ChangeCompany).Methods("PUT")
}
