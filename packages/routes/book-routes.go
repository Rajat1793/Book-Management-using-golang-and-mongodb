package routes

import (
	"fmt"

	"github.com/gorilla/mux"
	"main.go/packages/controllers"
)

var RegisterBookStoreRoutes = func(router *mux.Router) {

	fmt.Println("Welcome to Book Store Routes")

	router.HandleFunc("/book/", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/book/", controllers.GetBook).Methods("GET")
	router.HandleFunc("/book/{bookId}", controllers.GetBookById).Methods("GET")
	router.HandleFunc("/book/{bookId}", controllers.UpdateBook).Methods("PUT")
	router.HandleFunc("/book/{bookId}", controllers.DeleteBook).Methods("DELETE")
}
