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
	router.HandleFunc("/colleague-manager/auth", h.Login).Methods("POST")
	//-----------------------COLLEAGUE ONLY
	router.Handle("/colleague-manager/add-colleague", h.AuthMiddleware(h.AddColleague(), "colleague"))
	router.Handle("/colleague-manager/block-colleague/{id}", h.AuthMiddleware(h.BlockColleague(), "colleague"))
	router.Handle("/colleague-manager/add-client", h.AuthMiddleware(h.AddClient(), "colleague"))
	router.Handle("/colleague-manager/show-client-info/{id}", h.AuthMiddleware(h.ShowClientInfo(), "colleague"))
	router.Handle("/colleague-manager/change-client/{id}", h.AuthMiddleware(h.ChangeClient(), "colleague"))
	router.Handle("/colleague-manager/show-clients-applications-info", h.AuthMiddleware(h.ShowClientApplicationsInfo(), "colleague"))
	router.Handle("/colleague-manager/change-application-status/{id}", h.AuthMiddleware(h.ChangeApplicationStatus(), "colleague"))
	router.Handle("/colleague-manager/add-company", h.AuthMiddleware(h.AddCompany(), "colleague"))
	router.Handle("/colleague-manager/add-client-into-company", h.AuthMiddleware(h.AddClientIntoCompany(), "colleague"))
	router.Handle("/colleague-manager/change-company/{id}", h.AuthMiddleware(h.ChangeCompany(), "colleague"))
	//-------------------------CLIENTS ONLY
	router.Handle("/colleague-manager/show-client-info", h.AuthMiddleware(h.ShowClientInfo(), "client"))
	router.Handle("/colleague-manager/show-client-company-info", h.AuthMiddleware(h.ShowClientCompanysInfo(), "client"))
	router.Handle("/colleague-manager/add-application", h.AuthMiddleware(h.AddApplication(), "client"))
}
