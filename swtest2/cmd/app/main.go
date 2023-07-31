package main

import (
	"log"
	"net/http"
	"swtest2/internal/database"
	"swtest2/internal/service"
	"swtest2/internal/transport/handlers"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func main() {
	base, err := database.Connect()
	if err != nil {
		panic(err)
	}

	r := mux.NewRouter()

	repoDB := database.NewRepoDB(base)
	services := service.NewService(repoDB)
	handlers := handlers.NewHandler(services)
	handlers.ColleagueManagerRoutes(r)

	http.Handle("/", r)
	server := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}
	log.Fatal(server.ListenAndServe())
}
