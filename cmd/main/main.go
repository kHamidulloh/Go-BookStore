package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kHamidulloh/Go-BookStore/pkg/routes"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:8000", r))
}
