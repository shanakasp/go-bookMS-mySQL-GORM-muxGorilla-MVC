package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/princesp/go-bookms/pkg/routes"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStore(r)
	http.Handle("/",r)
	log.Fatal(http.ListenAndServe("Localhost:8000",r))
}