package routes

import (
	"github.com/gorilla/mux"
)

var RegisterBookStore = func (r *mux.Router){
	
	r.HandleFunc("/book/", controllers.PostsCreate).Methods("POST")
}