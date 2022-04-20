package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/invalder/fcc-go-11-projects/go-mysql-book-management/routes"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:9010", r))
}
