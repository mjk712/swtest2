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
			w.Write([]byte("error in add company:" + err.Error()))
			return
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
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("error in change company:" + err.Error()))
			return
		}
		res, _ := json.Marshal(company)
		w.WriteHeader(http.StatusCreated)
		w.Write(res)
	})
}

func (h *handler) AddClientIntoCompany() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		type clientCompany struct {
			ClientFio   string `json:"client_fio"`
			CompanyName string `json:"company_name"`
		}
		clientCompanyInfo := &clientCompany{}
		utils.ParseBody(r, clientCompanyInfo)
		err := h.service.AddClientIntoCompany(clientCompanyInfo.ClientFio, clientCompanyInfo.CompanyName)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("error in add client into company:" + err.Error()))
			return
		}
		f := fmt.Sprintf("Added %s into %s company.", clientCompanyInfo.ClientFio, clientCompanyInfo.CompanyName)
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
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("error in show client companys info:" + err.Error()))
			return
		}

		res, _ := json.Marshal(clientCompanysInfo)
		w.WriteHeader(http.StatusOK)
		w.Write(res)
	})
}
