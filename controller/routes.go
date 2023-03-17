package controller

import (
	"github.com/gorilla/mux"
	"net/http"
)

func Start(controller *Controller) error {
	router := mux.NewRouter()
	// get list of all books and create book
	router.HandleFunc("/books", controller.HandleBooks)
	// get, update, delete book information
	router.HandleFunc("/books/{id:[0-9]+}", controller.HandleBookById)
	// search book by title
	//router.HandleFunc("/books?{title}", controller.HandleSearch).Methods("GET")
	// get books in asc or desc order
	router.HandleFunc("/books/{order}", controller.HandleOrder).Methods("GET")
	return http.ListenAndServe(":8080", router)
}
