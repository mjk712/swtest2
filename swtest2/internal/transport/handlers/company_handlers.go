package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"swtest2/internal/models"
	"swtest2/internal/utils"

	"github.com/gorilla/mux"
)

func (h *handler) AddCompany() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		company := &models.Company{}
		utils.ParseBody(r, company)
		s, err := h.service.InsertCompany(company)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Error in add company"))
		}
		res, _ := json.Marshal(s)
		w.WriteHeader(http.StatusCreated)
		w.Write(res)
	})
}

func (h *handler) ChangeCompany() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		vars := mux.Vars(r)
		companyId := vars["id"]
		company := &models.Company{}
		utils.ParseBody(r, company)
		err := h.service.ChangeCompany(company, companyId)
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Shit error in change company"))
		}
		res, _ := json.Marshal(company)
		w.WriteHeader(http.StatusCreated)
		w.Write(res)
	})
}

func (h *handler) AddClientIntoCompany() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		type client_company struct {
			Client_fio   string `json:"client_fio"`
			Company_name string `json:"company_name"`
		}
		clientCompanyInfo := &client_company{}
		utils.ParseBody(r, clientCompanyInfo)
		err := h.service.AddClientIntoCompany(clientCompanyInfo.Client_fio, clientCompanyInfo.Company_name)
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Shit error in add client into company"))
		}
		f := fmt.Sprintf("Added %s into %s company.", clientCompanyInfo.Client_fio, clientCompanyInfo.Company_name)
		res, _ := json.Marshal(f)
		w.WriteHeader(http.StatusCreated)
		w.Write(res)
	})
}

func (h *handler) ShowClientCompanysInfo() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id, _ := r.Context().Value("clientId").(uint64)
		clientCompanysInfo, err := h.service.ShowClientCompanysInfo(strconv.FormatUint(id, 10))
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Shit error in show client companys info"))
		}

		res, _ := json.Marshal(clientCompanysInfo)
		w.WriteHeader(http.StatusOK)
		w.Write(res)
	})
}
