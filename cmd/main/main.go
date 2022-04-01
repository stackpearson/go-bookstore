package main

//all we're going to do here is create the server and tell bookstore-routes where our routes are

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	//this is our routes
	"github.com/sawyer/go-bookstore/pkg/routes"
)

func main() {
	r := mux.NewRouter()
	//this passes r to our RegisterBookStoreRoutes function in the bookstore-routes.go file
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	//this helps us listen & create a server at the specified address, it will log errors as configured
	log.Fatal(http.ListenAndServe("localhost:9010", r))
}
