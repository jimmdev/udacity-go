package main

import (
	"encoding/json"
	"testing"
)

// TestCreateACustomer checks if marshalling a customer to the json model works
func TestCreateACustomer(t *testing.T) {
	customerJimData := []byte(`{
		"name": "Jim Morrison",
		"role": "VIP customer",
		"email": "j.morrison@mail.com",
		"phone": 546879,
		"contacted": true
	}`)
	var customerJim Customer
	err := json.Unmarshal(customerJimData, &customerJim)
	if err != nil {
		t.Errorf("create customers returned an error: %v", err.Error())
	}
}
