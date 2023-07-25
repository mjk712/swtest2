package main

import (
	"log"
	"net/http"
	"swtest2/internal/transport/routes"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func main() {
	r := mux.NewRouter()
	routes.ColleagueManagerRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:8080", r))
}
