package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	// load some initial values to the customer DB
	loadInitialDatabaseValues()
	// set up the router and connect all endpoints
	router := mux.NewRouter().StrictSlash(true)
	// handle endpoint to get all customers
	router.HandleFunc("/customers", getCustomers).Methods("GET")
	// handle endpoint to get a single customer by id
	router.HandleFunc("/customers/{id}", getCustomer).Methods("GET")
	// handle endpoint to add a new customer
	router.HandleFunc("/customers", addCustomer).Methods("POST")
	// handle endpoint to update a customer with an ID
	router.HandleFunc("/customers/{id}", updateCustomer).Methods("PUT")
	// handle endpoint to delete a customer by ID
	router.HandleFunc("/customers/{id}", deleteCustomer).Methods("DELETE")
	// give a static response on the home path to show documentation
	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./public/")))
	// fire up the server
	fmt.Printf("Starting server on port :8080")
	http.ListenAndServe(":8080", router)
}
