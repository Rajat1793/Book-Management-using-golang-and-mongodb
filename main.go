package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"main.go/packages/routes"
)

func main() {
	fmt.Println("Welcome to BookStore Managemnet")
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:9010", r))
}
