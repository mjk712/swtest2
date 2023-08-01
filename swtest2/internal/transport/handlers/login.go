package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"swtest2/internal/models"
	"swtest2/internal/utils"
)

func (h *handler) Login(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		user := &models.User{}
		utils.ParseBody(r, user)
		token, err := h.service.GetUserToken(user.Username, user.Password)
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("error in login"))
			return
		}

		f := fmt.Sprintf("Succesfull login, your token - %s", token)
		res, _ := json.Marshal(f)
		w.WriteHeader(http.StatusOK)
		w.Write(res)

	case "GET":
		fmt.Fprintf(w, "only POST methods is allowed.")
	}
}
