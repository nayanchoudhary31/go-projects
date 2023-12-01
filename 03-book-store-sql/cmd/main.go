package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/nayanchoudhary31/03-book-store-sql/pkg/routes"
)

func main() {

	r := mux.NewRouter()
	routes.BookStoreRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:8010", r))
}
