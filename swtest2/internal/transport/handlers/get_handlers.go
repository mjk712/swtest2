package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"swtest2/internal/models"
	"swtest2/internal/service"
	"swtest2/internal/utils"

	"github.com/gorilla/mux"
)

func AddColleague(w http.ResponseWriter, r *http.Request) {

	colleague := &models.Colleague{}
	utils.ParseBody(r, colleague)
	err := service.AddColleague(colleague)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Shit error in add colleague"))
	}

	f := fmt.Sprintf("Created Colleague-%s", colleague.Fio)
	res, _ := json.Marshal(f)
	w.WriteHeader(http.StatusCreated)
	w.Write(res)
}

func BlockColleague(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	colleagueId := vars["id"]
	err := service.BlockColleague(colleagueId)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("OMG error in Block colleague"))
	}
	m := fmt.Sprintf("Block Colleague with id %s .", colleagueId)
	res, _ := json.Marshal(m)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func AddClient(w http.ResponseWriter, r *http.Request) {

	client := &models.Client{}
	utils.ParseBody(r, client)
	err := service.AddClient(client)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Shit error in add client"))
	}
	f := fmt.Sprintf("Created Client-%s", client.Fio)
	res, _ := json.Marshal(f)
	w.WriteHeader(http.StatusCreated)
	w.Write(res)
}

func ShowClientInfo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	clientId := vars["id"]
	clientInfo, err := service.ShowClientInfo(clientId)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("OMG error in ShowClientInfo"))
	}

	type client_info struct {
		Client_full_name       string
		Client_passport        string
		Client_email_telephone string
		Vacations              []string
	}

	var client_inf client_info

	for _, v := range clientInfo {
		client_inf.Client_email_telephone = v.Client_email_telephone
		client_inf.Client_passport = v.Client_passport
		client_inf.Client_full_name = v.Client_full_name
		client_inf.Vacations = append(client_inf.Vacations, v.Vacation_name)
	}

	res, _ := json.Marshal(client_inf)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func ChangeClient(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	clientId := vars["id"]
	client := &models.Client{}
	utils.ParseBody(r, client)
	err := service.ChangeClient(client, clientId)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Shit error in change client"))
	}
	f := fmt.Sprintf("Successfully change of the client-%s", client.Fio)
	res, _ := json.Marshal(f)
	w.WriteHeader(http.StatusCreated)
	w.Write(res)
}

func ShowClientApplicationsInfo(w http.ResponseWriter, r *http.Request) {

	clientsApplicationsInfo, err := service.ShowClientsApplicationsInfo()
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Shit error in show clients applications info"))
	}

	res, _ := json.Marshal(clientsApplicationsInfo)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func ChangeApplicationStatus(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	application_id := vars["id"]
	type application_status struct {
		Status string `json:"status"`
	}
	appstatus := &application_status{}
	utils.ParseBody(r, appstatus)
	err := service.ChangeApplicationStatus(application_id, appstatus.Status)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Shit error in change application status"))
	}
	f := fmt.Sprintf("Change application with id %s, status = %s", application_id, appstatus.Status)
	res, _ := json.Marshal(f)
	w.WriteHeader(http.StatusCreated)
	w.Write(res)
}

func AddCompany(w http.ResponseWriter, r *http.Request) {
	company := &models.Company{}
	utils.ParseBody(r, company)
	s, err := service.InsertCompany(company)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Error in add company"))
	}
	res, _ := json.Marshal(s)
	w.WriteHeader(http.StatusCreated)
	w.Write(res)
}

func AddClientIntoCompany(w http.ResponseWriter, r *http.Request) {

	type client_company struct {
		Client_fio   string `json:"client_fio"`
		Company_name string `json:"company_name"`
	}
	clientCompanyInfo := &client_company{}
	utils.ParseBody(r, clientCompanyInfo)
	err := service.AddClientIntoCompany(clientCompanyInfo.Client_fio, clientCompanyInfo.Company_name)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Shit error in add client into company"))
	}
	f := fmt.Sprintf("Added %s into %s company.", clientCompanyInfo.Client_fio, clientCompanyInfo.Company_name)
	res, _ := json.Marshal(f)
	w.WriteHeader(http.StatusCreated)
	w.Write(res)
}

func ChangeCompany(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	companyId := vars["id"]
	company := &models.Company{}
	utils.ParseBody(r, company)
	err := service.ChangeCompany(company, companyId)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Shit error in change company"))
	}
	f := fmt.Sprintf("Successfully change of the company-%s", company.Name)
	res, _ := json.Marshal(f)
	w.WriteHeader(http.StatusCreated)
	w.Write(res)
}
