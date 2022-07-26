package main

import (
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialetcs/mysql"
	"github.com/thneves/go-bookstore/pkg/routes"
	"log"
	"net/http"
)

func main() {
	route := mux.NewRouter()
	routes.RegisterBooksStoreRoutes(route)
	http.Handle("/", route)
	log.Fatal(http.ListenAndServe("localhost:8080", route))
}
