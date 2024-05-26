package routes

import (
	"github.com/gorilla/mux"
)

var RegisterBookStore = func (r *mux.Router){
	
	r.HandleFunc("/book/", controllers.PostsCreate).Methods("POST")
	r.HandleFunc("/book/", controllers.GetBook).Methods("GET")
	r.HandleFunc("/book/{bookid}", controllers.GetBookById).Methods("GET")
	r.HandleFunc("/book/{bookid}", controllers.UpdateBook).Methods("PUT")
	r.HandleFunc("/book/{bookid}", controllers.DeleteBook).Methods("DELETE")


}