package routes

import (
	controllers "gorilla_mysql/pkg/controllers"

	"github.com/gorilla/mux"
)

var RegisterBookStoreRoutes = func(r *mux.Router) {
	r.HandleFunc("/book", controllers.CreateBook).Methods("POST")
	r.HandleFunc("/book", controllers.GetBook).Methods("GET")
	r.HandleFunc("/book", controllers.GetBookById).Methods("GET")
	r.HandleFunc("/book", controllers.UpdateBook).Methods("PUT")
	r.HandleFunc("/book", controllers.DeleteBook).Methods("DELETE")

}
