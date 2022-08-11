package main

import (
	"encoding/json"
	"testing"
)

// initially fills the database with some content
var customerIds = loadInitialDatabaseValues()

// and create one more user to test with
var extraCustomer = createJim()

// TestAddCustomer2DB adds a customer to DB
// test is regarded successful if a proper ID is assigned
func TestAddCustomer2DB(t *testing.T) {
	returnValue := addCustomerToDB(extraCustomer)
	if len(returnValue.Id) == 0 {
		t.Errorf("creating customer did not return a proper id")
	}
}

// TestGetCustomerFromDB tries to get a user from the DB
// test is regarded successful if GET returns success
// and the IDs of the requested customer matches the request-id
func TestGetCustomerFromDB(t *testing.T) {
	fetchedCustomer, success := getCustomerFromDB(customerIds[0])
	if !success {
		t.Errorf("GET customer from DB was not successful")
	}
	if customerIds[0] != fetchedCustomer.Id {
		t.Errorf("GET customer returned the wrong customer")
	}
}

// TestUpdateCustomerInDB gets a customer from DB
// then updates a value and tries to update the user in the DB
// test is regarded successful if UPDATE returns success
// and the return value contains the updated value
func TestUpdateCustomerInDB(t *testing.T) {
	customerZero, _ := getCustomerFromDB(customerIds[0])
	customerZero.Phone = 666666
	updatedCustomer, success := updateCustomerInDB(customerIds[0], customerZero)
	if !success {
		t.Errorf("UPDATE customer from DB was not successful")
	}
	if updatedCustomer.Phone != 666666 {
		t.Errorf("UPDATE customer returned the wrong phone number")
	}
}

// TestDeleteCustomerFromDB tries to delete a user in the DB
// test is regarded successful if DELETE returns success
// and a GET customer cannot find the customer anymore
func TestDeleteCustomerFromDB(t *testing.T) {
	success := deleteCustomerFromDB(customerIds[0])
	if !success {
		t.Errorf("DELETE customer from DB was not successful")
	}
	_, found := getCustomerFromDB(customerIds[0])
	if found {
		t.Errorf("Customer does still exist after deletion")
	}
}

// TestGetAllCustomers2DB tries to get all customers from DB
// test is regarded successful if GET all returns two or more customers
func TestGetAllCustomers2DB(t *testing.T) {
	allCustomers := getAllCustomersFromDB()
	if len(allCustomers) < 2 {
		t.Errorf("GET all customers from DB did not return the amount of customers")
	}
}

// helps to create another customer
func createJim() Customer {
	data := []byte(`{
		"name": "Jim Morrison",
		"role": "VIP customer",
		"email": "j.morrison@mail.com",
		"phone": 546879,
		"contacted": true
	}`)
	var customerJim Customer
	json.Unmarshal(data, &customerJim)
	return customerJim
}
