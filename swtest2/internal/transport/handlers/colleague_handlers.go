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

func AddColleague() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

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
	})
}

func BlockColleague() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
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
	})
}
