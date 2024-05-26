package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	initializers "github.com/princesp/go-bookms/inilializers"
	"github.com/princesp/go-bookms/routes"
)

func init(){
	initializers.LoadEnvVariables()
}
func main() {
	r := mux.NewRouter()
	routes.RegisterBookStore(r)
	http.Handle("/",r)
	log.Fatal(http.ListenAndServe("Localhost:8000",r))
}