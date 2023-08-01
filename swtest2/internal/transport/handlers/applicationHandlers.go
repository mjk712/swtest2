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
			w.Write([]byte("error in add application"))
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
		applicationId := vars["id"]
		type applicationStatus struct {
			Status string `json:"status"`
		}
		appstatus := &applicationStatus{}
		utils.ParseBody(r, appstatus)
		err := h.service.ChangeApplicationStatus(applicationId, appstatus.Status)
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("error in change application status"))
			return
		}
		f := fmt.Sprintf("Change application with id %s, status = %s", applicationId, appstatus.Status)
		res, _ := json.Marshal(f)
		w.WriteHeader(http.StatusCreated)
		w.Write(res)
	})
}
