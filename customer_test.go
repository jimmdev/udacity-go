package main

import (
	"encoding/json"
	"testing"
)

func TestCreateACustomer(t *testing.T) {
	customerJimData := []byte(`{
		"id": "b8451078-17fa-11ed-861d-0242ac120002",
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
