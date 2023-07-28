package routes

import (
	"swtest2/internal/transport/handlers"

	"github.com/gorilla/mux"
)

var ColleagueManagerRoutes = func(router *mux.Router) {
	//LOGIN IN
	router.HandleFunc("/colleaguemanager/auth", handlers.Login).Methods("POST")
	//-----------------------COLLEAGUE ONLY
	router.Handle("/colleaguemanager/addcolleague", handlers.AuthMiddleware(handlers.AddColleague(), "colleague"))
	router.Handle("/colleaguemanager/blockcolleague/{id}", handlers.AuthMiddleware(handlers.BlockColleague(), "colleague"))
	router.Handle("/colleaguemanager/addclient", handlers.AuthMiddleware(handlers.AddClient(), "colleague"))
	router.Handle("/colleaguemanager/showclientinfo/{id}", handlers.AuthMiddleware(handlers.ShowClientInfo(), "colleague"))
	router.Handle("/colleaguemanager/changeclient/{id}", handlers.AuthMiddleware(handlers.ChangeClient(), "colleague"))
	router.Handle("/colleaguemanager/showclientsapplicationsinfo", handlers.AuthMiddleware(handlers.ShowClientApplicationsInfo(), "colleague"))
	router.Handle("/colleaguemanager/changeapplicationstatus/{id}", handlers.AuthMiddleware(handlers.ChangeApplicationStatus(), "colleague"))
	router.Handle("/colleaguemanager/addcompany", handlers.AuthMiddleware(handlers.AddCompany(), "colleague"))
	router.Handle("/colleaguemanager/addclientintocompany", handlers.AuthMiddleware(handlers.AddClientIntoCompany(), "colleague"))
	router.Handle("/colleaguemanager/changecompany/{id}", handlers.AuthMiddleware(handlers.ChangeCompany(), "colleague"))
	//-------------------------CLIENTS ONLY
	router.Handle("/colleaguemanager/showclientinfo", handlers.AuthMiddleware(handlers.ShowClientInfo(), "client"))
	router.Handle("/colleaguemanager/showclientcompanyinfo", handlers.AuthMiddleware(handlers.ShowClientCompanysInfo(), "client"))
	router.Handle("/colleaguemanager/addapplication", handlers.AuthMiddleware(handlers.AddApplication(), "client"))
}
