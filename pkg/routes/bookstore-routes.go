package routes

import (
	"github.com/cbondoc/go-bookstore/pkg/controllers"
	"github.com/gorilla/mux"
)


var RegisterBookStoreRoutes = func(router *mux.Router) {
	router.Handle("/book/", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/book/", controllers.GetBook).Methods("GET")
	router.HandleFunc("/book/{bookId}", controllers.GetBookById).Methods("GET")
	router.HandleFunc("/book/{bookId}", controllers.UpdateBook).Methods("PUT")
	router.HandleFunc("/book/{bookId}", controllers.DeleteBook).Methods("DELETE")
}
