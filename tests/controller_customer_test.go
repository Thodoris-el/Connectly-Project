package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	entity "github.com/Thodoris-el/Coonectly-Project/api/Entity"
	"github.com/gorilla/mux"
)

func TestCreateCustomer(t *testing.T) {

	refreshTables()

	samples := []string{
		`{"first_name": "John","last_name": "Dir","facebook_id": "6706612322695175","sent_message": true}`,
		`{"first_name": "Nick","last_name": "Dir","facebook_id": "8706612322695175","sent_message": true}`,
		`{"first_name": "Theo","last_name": "Dir","facebook_id": "9806612322695175","sent_message": true}`,
		`{"first_name": "Bo","last_name": "Dir","facebook_id": "6798612322695175","sent_message": true}`,
	}
	answers := []string{
		"John",
		"Nick",
		"Theo",
		"Bo",
	}

	for i, tmp := range samples {
		req, err := http.NewRequest("POST", "/customer", bytes.NewBufferString(tmp))
		if err != nil {
			t.Errorf("error: %v\n", err)
		}
		recorded := httptest.NewRecorder()
		handler := http.HandlerFunc(server.CreateCustomer)
		handler.ServeHTTP(recorded, req)

		resp := make(map[string]interface{})
		err = json.Unmarshal(recorded.Body.Bytes(), &resp)
		if err != nil {
			t.Errorf("cant convert json: %v\n", err)
		}
		if resp["first_name"] != answers[i] {
			t.Errorf("wrong values")
		}
	}
}

func TestGetCustomers(t *testing.T) {

	refreshTables()

	_, err := createTwoCustomers()
	if err != nil {
		t.Errorf("error creating customers")
	}

	req, err := http.NewRequest("GET", "/customer", nil)
	if err != nil {
		t.Errorf("error: %v\n", err)
	}
	recorded := httptest.NewRecorder()
	handler := http.HandlerFunc(server.GetCustomers)
	handler.ServeHTTP(recorded, req)

	var customers []entity.Customer
	err = json.Unmarshal(recorded.Body.Bytes(), &customers)
	if err != nil {
		t.Errorf("error while unmarshal")
	}
	if len(customers) != 2 {
		t.Errorf("wrong number of customers")
	}
}

func TestGetCustomerById(t *testing.T) {

	refreshTables()

	_, err := createACustomer()
	if err != nil {
		t.Errorf("error creating a customer")
	}

	req, err := http.NewRequest("GET", "/customer", nil)
	if err != nil {
		t.Errorf("error: %v\n", err)
	}
	req = mux.SetURLVars(req, map[string]string{"id": "1"})
	recorded := httptest.NewRecorder()
	handler := http.HandlerFunc(server.GetCustomerById)
	handler.ServeHTTP(recorded, req)

	responseMap := make(map[string]interface{})
	err = json.Unmarshal(recorded.Body.Bytes(), &responseMap)
	if err != nil {
		t.Errorf("error while unmarshal")
	}
	if responseMap["first_name"] != "John" {
		t.Errorf("wrong value")
	}
}

func TestUpdateCustomerC(t *testing.T) {

	refreshTables()

	_, err := createACustomer()
	if err != nil {
		t.Errorf("error creating a customer")
	}
	sample := `{"first_name": "Nick"}`
	req, err := http.NewRequest("PUT", "/customer", bytes.NewBufferString(sample))
	if err != nil {
		t.Errorf("error: %v\n", err)
	}

	req = mux.SetURLVars(req, map[string]string{"id": "1"})
	recorded := httptest.NewRecorder()
	handler := http.HandlerFunc(server.UpdateCustomer)
	handler.ServeHTTP(recorded, req)

	responseMap := make(map[string]interface{})
	err = json.Unmarshal(recorded.Body.Bytes(), &responseMap)
	if err != nil {
		t.Errorf("error while unmarshal")
	}
	if responseMap["first_name"] != "Nick" {
		t.Errorf("wrong value")
	}
}

func TestDeleteCustomersc(t *testing.T) {

	refreshTables()

	_, err := createACustomer()
	if err != nil {
		t.Errorf("error creating a customer")
	}

	req, err := http.NewRequest("DELETE", "/customer", nil)
	if err != nil {
		t.Errorf("error: %v\n", err)
	}
	req = mux.SetURLVars(req, map[string]string{"id": "1"})
	recorded := httptest.NewRecorder()
	handler := http.HandlerFunc(server.DeleteCustomer)
	handler.ServeHTTP(recorded, req)

	var customers entity.Customer
	getC, err := customers.FindAllCustomers(server.DB)
	if err != nil {
		t.Errorf("error while getting customers")
	}
	if len(*getC) != 0 {
		t.Errorf("wrong number of customers")
	}
}
