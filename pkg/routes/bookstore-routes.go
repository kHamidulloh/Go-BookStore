package routes

import (
	"github.com/gorilla/mux"
	"github.com/kHamidulloh/Go-BookStore/pkg/controllers"
)

var RegisterBookStoreRoutes = func(r *mux.Router) {
	r.HandleFunc("/book/", controllers.CreateBook).Methods("POST")
	r.HandleFunc("/book/", controllers.GetBook).Methods("GET")
	r.HandleFunc("/book/{bookId}", controllers.GetById).Methods("GET")
	r.HandleFunc("/book/{bookId}", controllers.UpgradeBook).Methods("PUT")
	r.HandleFunc("/book/{bookId}", controllers.DeleteBook).Methods("DELETE")
}
