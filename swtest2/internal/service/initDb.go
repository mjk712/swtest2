package service

import (
	"fmt"
	"net/http"
	"swtest2/internal/database"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var db *sqlx.DB

func init() {
	err := database.Connect()
	if err != nil {
		fmt.Println("Bad DataBase")
		fmt.Println(http.StatusBadRequest)
		panic(err)
	}
	db = database.GetDB()
}
