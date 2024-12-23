package main

import (
	"GO-BOOKSTORE/pkg/routes"
	"log"
	"net/http"

	"github.com/go-gorm/gorm/pkg/routes"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:9010", r))

}
