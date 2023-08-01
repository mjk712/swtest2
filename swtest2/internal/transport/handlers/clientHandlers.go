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

func (h *handler) AddClient() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		client := &models.Client{}
		utils.ParseBody(r, client)
		err := h.service.AddClient(client)
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("error in add client"))
			return
		}
		f := fmt.Sprintf("Created Client-%s", client.Fio)
		res, _ := json.Marshal(f)
		w.WriteHeader(http.StatusCreated)
		w.Write(res)
	})
}

func (h *handler) ShowClientInfo() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		clientId, ok := vars["id"]
		if !ok {
			id := r.Context().Value("clientId").(uint64)
			clientId = strconv.FormatUint(id, 10)
		}
		clientInfo, err := h.service.ShowClientInfo(clientId)
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("error in ShowClientInfo"))
			return
		}

		type clientInformation struct {
			ClientFullname       string
			ClientPassport       string
			ClientEmailTelephone string
			Vacations            []string
		}

		var clientInf clientInformation

		for _, v := range clientInfo {
			clientInf.ClientEmailTelephone = v.ClientEmailTelephone
			clientInf.ClientPassport = v.ClientPassport
			clientInf.ClientFullname = v.ClientFullname
			clientInf.Vacations = append(clientInf.Vacations, v.VacationName)
		}

		res, _ := json.Marshal(clientInf)
		w.WriteHeader(http.StatusOK)
		w.Write(res)
	})
}

func (h *handler) ChangeClient() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		vars := mux.Vars(r)
		clientId := vars["id"]
		client := &models.Client{}
		utils.ParseBody(r, client)
		err := h.service.ChangeClient(client, clientId)
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("error in change client"))
			return
		}
		res, _ := json.Marshal(client)
		w.WriteHeader(http.StatusCreated)
		w.Write(res)
	})
}

func (h *handler) ShowClientApplicationsInfo() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		clientsApplicationsInfo, err := h.service.ShowClientsApplicationsInfo()
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("error in show clients applications info"))
			return
		}

		res, _ := json.Marshal(clientsApplicationsInfo)
		w.WriteHeader(http.StatusOK)
		w.Write(res)
	})
}
