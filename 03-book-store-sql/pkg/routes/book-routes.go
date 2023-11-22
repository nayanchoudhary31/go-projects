package routes

import (
	"github.com/gorilla/mux"
	"github.com/nayanchoudhary31/03-book-store-sql/pkg/controller"
)

var BookStoreRoutes = func(router *mux.Router) {
	router.HandleFunc("/books/", controller.CreateBook).Methods("POST")

}
