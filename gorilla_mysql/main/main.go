package main

import (
	router "gorilla_mysql/pkg/routes"
	"log"

	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	router.RegisterBookStoreRoutes(r)
	r.Headers("Content-Type", "application/json")

	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:8080", r))
}
