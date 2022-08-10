package main

import (
	"encoding/json"
	"testing"
)

var customerJimData = []byte(`{
		"name": "Jim Morrison",
		"role": "VIP customer",
		"email": "j.morrison@mail.com",
		"phone": 546879,
		"contacted": true
	}`)

var customerJohnData = []byte(`{
		"name": "John Densmore",
		"role": "Frequent customer",
		"email": "j.densmore@online.com",
		"phone": 213546,
		"contacted": false
	}`)

// TestAddCustomer2DB creates a customer and stores it to DB
// test is regarded successful if a proper ID is assigned
func TestAddCustomer2DB(t *testing.T) {
	var customerJim Customer
	err := json.Unmarshal(customerJimData, &customerJim)
	if err != nil {
		t.Errorf("parsing Jim returned an error: %v", err.Error())
	}
	returnValue := addCustomerToDB(customerJim)
	if len(returnValue.Id) == 0 {
		t.Errorf("creating customer did not return a proper id")
	}
}

// TestGetCustomerFromDB creates a customer and stores it to DB
// then tries to get the user from the DB
// test is regarded successful if GET returns success
// and the IDs of the requested customer matches the request-id
func TestGetCustomerFromDB(t *testing.T) {
	var customerJim Customer
	err := json.Unmarshal(customerJimData, &customerJim)
	if err != nil {
		t.Errorf("parsing Jim returned an error: %v", err.Error())
	}
	returnValue := addCustomerToDB(customerJim)
	if len(returnValue.Id) == 0 {
		t.Errorf("creating customer did not return a proper id")
	}
	fetchedCustomer, succ := getCustomerFromDB(returnValue.Id)
	if !succ {
		t.Errorf("GET customer from DB was not successful")
	}
	if returnValue.Id != fetchedCustomer.Id {
		t.Errorf("GET customer returned the wrong customer")
	}
}

// TestUpdateCustomerInDB creates a customer and stores it to DB
// then updates a value and tries to update the user in the DB
// test is regarded successful if UPDATE returns success
// and the return value contains the updated value
func TestUpdateCustomerInDB(t *testing.T) {
	var customerJim Customer
	err := json.Unmarshal(customerJimData, &customerJim)
	if err != nil {
		t.Errorf("parsing Jim returned an error: %v", err.Error())
	}
	returnValue := addCustomerToDB(customerJim)
	if len(returnValue.Id) == 0 {
		t.Errorf("creating customer did not return a proper id")
	}
	customerJim.Phone = 879546
	updatedCustomer, succ := updateCustomerInDB(returnValue.Id, customerJim)
	if !succ {
		t.Errorf("UPDATE customer from DB was not successful")
	}
	if updatedCustomer.Phone != 879546 {
		t.Errorf("UPDATE customer returned the wrong phone number")
	}
}

// TestDeleteCustomerFromDB creates a customer and stores it to DB
// then tries to delete the user in the DB again
// test is regarded successful if DELETE returns success
// and a GET customer cannot find the customer anymore
func TestDeleteCustomerFromDB(t *testing.T) {
	var customerJim Customer
	err := json.Unmarshal(customerJimData, &customerJim)
	if err != nil {
		t.Errorf("parsing Jim returned an error: %v", err.Error())
	}
	returnValue := addCustomerToDB(customerJim)
	if len(returnValue.Id) == 0 {
		t.Errorf("creating customer did not return a proper id")
	}
	succ := deleteCustomerFromDB(returnValue.Id)
	if !succ {
		t.Errorf("DELETE customer from DB was not successful")
	}
	_, found := getCustomerFromDB(returnValue.Id)
	if found {
		t.Errorf("Customer does still exist after deletion")
	}
}

// TestGetAllCustomers2DB creates two customers and stores them to DB
// then tries to get all customers from DB
// test is regarded successful if GET all returns two or more customers
// (DB was not cleaned initially)
func TestGetAllCustomers2DB(t *testing.T) {
	var customerJim Customer
	err := json.Unmarshal(customerJimData, &customerJim)
	if err != nil {
		t.Errorf("parsing Jim returned an error: %v", err.Error())
	}
	var customerJohn Customer
	err2 := json.Unmarshal(customerJohnData, &customerJohn)
	if err2 != nil {
		t.Errorf("parsing John returned an error: %v", err.Error())
	}
	addCustomerToDB(customerJim)
	addCustomerToDB(customerJohn)
	allCustomers := getAllCustomersFromDB()
	if len(allCustomers) < 2 {
		t.Errorf("GET all customers from DB did not return the amount of customers")
	}
}
