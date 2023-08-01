package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"swtest2/internal/models"
	"swtest2/internal/utils"

	"github.com/gorilla/mux"
)

func (h *handler) AddColleague() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		colleague := &models.Colleague{}
		utils.ParseBody(r, colleague)
		err := h.service.AddColleague(colleague)
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("error in add colleague"))
			return
		}

		f := fmt.Sprintf("Created Colleague-%s", colleague.Fio)
		res, _ := json.Marshal(f)
		w.WriteHeader(http.StatusCreated)
		w.Write(res)
	})
}

func (h *handler) BlockColleague() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		colleagueId := vars["id"]
		err := h.service.BlockColleague(colleagueId)
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("error in Block colleague"))
			return
		}
		m := fmt.Sprintf("Block Colleague with id %s .", colleagueId)
		res, _ := json.Marshal(m)
		w.WriteHeader(http.StatusOK)
		w.Write(res)
	})
}
