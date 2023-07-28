package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"swtest2/internal/models"
	"swtest2/internal/service"
	"swtest2/internal/utils"

	"github.com/gorilla/mux"
)

func AddClient() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

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
	})
}

func ShowClientInfo() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		clientId, ok := vars["id"]
		if !ok {
			id := r.Context().Value("clientId").(uint64)
			clientId = strconv.FormatUint(id, 10)
		}
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
	})
}

func ChangeClient() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

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
		res, _ := json.Marshal(client)
		w.WriteHeader(http.StatusCreated)
		w.Write(res)
	})
}

func ShowClientApplicationsInfo() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		clientsApplicationsInfo, err := service.ShowClientsApplicationsInfo()
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Shit error in show clients applications info"))
		}

		res, _ := json.Marshal(clientsApplicationsInfo)
		w.WriteHeader(http.StatusOK)
		w.Write(res)
	})
}
