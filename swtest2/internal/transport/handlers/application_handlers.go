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

func (h *handler) AddApplication() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id, _ := r.Context().Value("clientId").(uint64)
		application := &models.Application{}
		utils.ParseBody(r, application)
		err := h.service.AddApplication(application, strconv.FormatUint(id, 10))
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("error in add application:" + err.Error()))
			return
		}
		res, _ := json.Marshal(application)
		w.WriteHeader(http.StatusCreated)
		w.Write(res)
	})
}

func (h *handler) ChangeApplicationStatus() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		application_id := vars["id"]
		type application_status struct {
			Status string `json:"status"`
		}
		appstatus := &application_status{}
		utils.ParseBody(r, appstatus)
		err := h.service.ChangeApplicationStatus(application_id, appstatus.Status)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("error in change application status" + err.Error()))
			return
		}
		f := fmt.Sprintf("Change application with id %s, status = %s", application_id, appstatus.Status)
		res, _ := json.Marshal(f)
		w.WriteHeader(http.StatusCreated)
		w.Write(res)
	})
}
