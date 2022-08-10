package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

// the handlers define the reactions to a call to the different endpoints

// getCustomers tries to fetch the customer with the given ID from the DB
// returns 200 and the customer as json if it was found
// returns 404 if no id was given or customer was not found
func getCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	customerId := mux.Vars(r)["id"]
	if len(customerId) == 0 {
		var customer Customer
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(customer)
	} else {
		foundCustomer, success := getCustomerFromDB(customerId)
		if success {
			w.WriteHeader(http.StatusOK)
		} else {
			w.WriteHeader(http.StatusNotFound)
		}
		json.NewEncoder(w).Encode(foundCustomer)
	}
}

// addCustomer tries to add the given customer to the DB
// returns 201 and the customer upon success
// return 400 if customer data to add was invalid
func addCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var customer Customer
	err := json.NewDecoder(r.Body).Decode(&customer)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(customer)
	} else {
		addedCustomer := addCustomerToDB(customer)
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(addedCustomer)
	}
}

// deleteCustomer tries to delete the customer with the given ID
// returns 200 if user was deleted
// returns 404 if no id was given or user could not be deleted
func deleteCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	customerId := mux.Vars(r)["id"]
	if len(customerId) == 0 {
		w.WriteHeader(http.StatusNotFound)
	} else {
		success := deleteCustomerFromDB(customerId)
		if success {
			w.WriteHeader(http.StatusOK)
		} else {
			w.WriteHeader(http.StatusNotFound)
		}
	}
	var customer Customer
	json.NewEncoder(w).Encode(customer)
}

// getCustomers fetches all customers from DB and returns the in a json array
func getCustomers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	customers := getAllCustomersFromDB()
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(customers)
}

// updateCustomer updated the customer with the given id with the given customer information
// returns 200 and the updated customer as json if update is successful
// returns 400 if ID not given or customer data not valid
// returns 404 if the ID could not be found in the database
func updateCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	customerId := mux.Vars(r)["id"]
	var customer Customer
	if len(customerId) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(customer)
	} else {
		err := json.NewDecoder(r.Body).Decode(&customer)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(customer)
		} else {
			updatedCustomer, success := updateCustomerInDB(customerId, customer)
			if !success {
				w.WriteHeader(http.StatusNotFound)
			} else {
				w.WriteHeader(http.StatusOK)
			}
			json.NewEncoder(w).Encode(updatedCustomer)
		}
	}
}
