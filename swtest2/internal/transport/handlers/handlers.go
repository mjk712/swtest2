package handlers

import (
	"swtest2/internal/service"

	"github.com/gorilla/mux"
)

type handler struct {
	service service.Service
}

type Handler interface {
	ColleagueManagerRoutes(router *mux.Router)
}

func NewHandler(service service.Service) Handler {
	return &handler{service}
}

func (h *handler) ColleagueManagerRoutes(router *mux.Router) {
	//LOGIN IN
	router.HandleFunc("/colleaguemanager/auth", h.Login).Methods("POST")
	//-----------------------COLLEAGUE ONLY
	router.Handle("/colleaguemanager/addcolleague", h.AuthMiddleware(h.AddColleague(), "colleague"))
	router.Handle("/colleaguemanager/blockcolleague/{id}", h.AuthMiddleware(h.BlockColleague(), "colleague"))
	router.Handle("/colleaguemanager/addclient", h.AuthMiddleware(h.AddClient(), "colleague"))
	router.Handle("/colleaguemanager/showclientinfo/{id}", h.AuthMiddleware(h.ShowClientInfo(), "colleague"))
	router.Handle("/colleaguemanager/changeclient/{id}", h.AuthMiddleware(h.ChangeClient(), "colleague"))
	router.Handle("/colleaguemanager/showclientsapplicationsinfo", h.AuthMiddleware(h.ShowClientApplicationsInfo(), "colleague"))
	router.Handle("/colleaguemanager/changeapplicationstatus/{id}", h.AuthMiddleware(h.ChangeApplicationStatus(), "colleague"))
	router.Handle("/colleaguemanager/addcompany", h.AuthMiddleware(h.AddCompany(), "colleague"))
	router.Handle("/colleaguemanager/addclientintocompany", h.AuthMiddleware(h.AddClientIntoCompany(), "colleague"))
	router.Handle("/colleaguemanager/changecompany/{id}", h.AuthMiddleware(h.ChangeCompany(), "colleague"))
	//-------------------------CLIENTS ONLY
	router.Handle("/colleaguemanager/showclientinfo", h.AuthMiddleware(h.ShowClientInfo(), "client"))
	router.Handle("/colleaguemanager/showclientcompanyinfo", h.AuthMiddleware(h.ShowClientCompanysInfo(), "client"))
	router.Handle("/colleaguemanager/addapplication", h.AuthMiddleware(h.AddApplication(), "client"))
}
