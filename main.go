package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	// set up the router and connect all endpoints
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/customers", getCustomers).Methods("GET")
	router.HandleFunc("/customers/{id}", getCustomer).Methods("GET")
	router.HandleFunc("/customers", addCustomer).Methods("POST")
	router.HandleFunc("/customers/{id}", updateCustomer).Methods("PUT")
	router.HandleFunc("/customer/{id}", deleteCustomer).Methods("DELETE")
	http.ListenAndServe(":8080", router)

}
