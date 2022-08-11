package main

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"os"
)

// this map functions as a DB for now
var customerMap = make(map[string]Customer)

// addCustomerToDB assigns an id and adds the customer to the database
// returns the customer with ID after it was added
func addCustomerToDB(customer Customer) Customer {
	customer.Id = uuid.New().String()
	customerMap[customer.Id] = customer
	return customer
}

// getCustomerFromDB looks for the customer with the given id in the database
// return the customer or an empty object and a boolean to indicate if it was found
func getCustomerFromDB(customerId string) (Customer, bool) {
	customerFromMap, found := customerMap[customerId]
	return customerFromMap, found
}

// updateCustomerInDB looks for the customer with the given id in the database
// returns false if it is not found, updates the customer if found
// returns true and the updated customer, after successful update
func updateCustomerInDB(customerId string, customer Customer) (Customer, bool) {
	customerFromMap, found := customerMap[customerId]
	if !found {
		return customerFromMap, found
	}
	customerMap[customerId] = customer
	return customerMap[customerId], true
}

// getAllCustomersFromDB returns a slice of all customers that are currently in the DB
func getAllCustomersFromDB() []Customer {
	var allCustomers []Customer
	for _, customer := range customerMap {
		allCustomers = append(allCustomers, customer)
	}
	return allCustomers
}

// deleteCustomerFromDB looks for the customer with the given id in the database
// returns false if it is not found, deletes the customer if found
// returns true after successful delete
func deleteCustomerFromDB(customerId string) bool {
	_, found := customerMap[customerId]
	if !found {
		return found
	}
	delete(customerMap, customerId)
	return true
}

// loadInitialDatabaseValues reads the customers.json
// and adds the content as initial values to the DB
func loadInitialDatabaseValues() []string {
	dat, err := os.ReadFile("./customers.json")
	var ids []string
	if err != nil {
		fmt.Print("Could not load initial data")
	} else {
		var customers []Customer
		json.Unmarshal(dat, &customers)
		for _, customer := range customers {
			addedCustomer := addCustomerToDB(customer)
			ids = append(ids, addedCustomer.Id)
		}
	}
	return ids
}
