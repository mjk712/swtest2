package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"swtest2/internal/models"
	"swtest2/internal/service"
	"swtest2/internal/utils"
)

func Login(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		user := &models.User{}
		utils.ParseBody(r, user)
		token, err := service.GetUserToken(user.Username, user.Password)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Shit error in login"))
		}

		f := fmt.Sprintf("Succesfull login, your token - %s", token)
		res, _ := json.Marshal(f)
		w.WriteHeader(http.StatusOK)
		w.Write(res)

	case "GET":
		fmt.Fprintf(w, "only POST methods is allowed.")
	}
}
