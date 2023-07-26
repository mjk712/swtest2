package routes

import (
	"swtest2/internal/transport/handlers"

	"github.com/gorilla/mux"
)

var ColleagueManagerRoutes = func(router *mux.Router) {
	//LOGIN IN
	router.HandleFunc("/colleaguemanager/auth", handlers.Login).Methods("POST")
	//-----------------------COLLEAGUE ONLY
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
	//router.HandleFunc("/colleaguemanager/showclientcompanyinfo", handlers.ShowClientCompanysInfo).Methods("GET")
	//router.HandleFunc("/colleaguemanager/addapplication", handlers.AddApplication).Methods("POST")
	//-------------------------CLIENTS ONLY
	router.Handle("/colleaguemanager/showclientinfo", handlers.AuthMiddleware(handlers.ShowClientCompanysInfo(), "client")) //СДЕЛАТЬ
	router.Handle("/colleaguemanager/showclientcompanyinfo", handlers.AuthMiddleware(handlers.ShowClientCompanysInfo(), "client"))
	router.Handle("/colleaguemanager/addapplication", handlers.AuthMiddleware(handlers.AddApplication(), "client"))
}
