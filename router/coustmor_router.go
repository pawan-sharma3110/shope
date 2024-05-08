package router

import (
	"shope/handler"

	"github.com/gorilla/mux"
)

func CustomerRouter() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/customer/create", handler.CreateCustomer).Methods("POST")
	router.HandleFunc("/customer/getall", handler.GetAllCustomer).Methods("GET")
	router.HandleFunc("/customer/searchbyid/{id}", handler.SearchById).Methods("GET")
	return router
}
