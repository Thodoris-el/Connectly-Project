package controllers

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"
	"time"

	entity "github.com/Thodoris-el/Coonectly-Project/api/Entity"
	"github.com/gorilla/mux"
)

//Create a customer
func (server *Server) CreateCustomer(resp http.ResponseWriter, request *http.Request) {

	body, err := io.ReadAll(request.Body)

	if err != nil {
		http.Error(resp, err.Error(), http.StatusBadRequest)
		return
	}

	customer := entity.Customer{}
	err = json.Unmarshal(body, &customer)
	if err != nil {
		http.Error(resp, err.Error(), http.StatusBadRequest)
		return
	}

	customer.CreatedAt = time.Now()
	customer.UpdatedAt = time.Now()
	customerCreated, err := customer.SaveCustomer(server.DB)

	if err != nil {
		http.Error(resp, err.Error(), http.StatusBadRequest)
		return
	}

	resp.Header().Set("Content-Type", "application/json")
	resp.WriteHeader(http.StatusOK)
	err = json.NewEncoder(resp).Encode(customerCreated)
	if err != nil {
		log.Println(err)
	}

}

//Get all customers
func (server *Server) GetCustomers(resp http.ResponseWriter, request *http.Request) {

	customer := entity.Customer{}
	customers, err := customer.FindAllCustomers(server.DB)

	if err != nil {
		http.Error(resp, err.Error(), http.StatusBadRequest)
		return
	}

	resp.Header().Set("Content-Type", "application/json")
	resp.WriteHeader(http.StatusOK)
	err = json.NewEncoder(resp).Encode(customers)
	if err != nil {
		log.Println(err)
	}
}

//Get Customer By Id
func (server *Server) GetCustomerById(resp http.ResponseWriter, request *http.Request) {

	vars := mux.Vars(request)
	C_id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		http.Error(resp, err.Error(), http.StatusBadRequest)
		return
	}

	customer := entity.Customer{}

	customerGet, err := customer.FindCustomerByID(server.DB, C_id)
	if err != nil {
		http.Error(resp, err.Error(), http.StatusBadRequest)
		return
	}

	resp.Header().Set("Content-Type", "application/json")
	resp.WriteHeader(http.StatusOK)
	err = json.NewEncoder(resp).Encode(customerGet)
	if err != nil {
		log.Println(err)
	}
}

//Get Customer By facebook id
func (server *Server) GetCustomerByFacebookId(resp http.ResponseWriter, request *http.Request) {

	vars := mux.Vars(request)
	C_id := vars["facebook_id"]

	customer := entity.Customer{}

	customerGet, err := customer.FindByFacebookId(server.DB, C_id)
	if err != nil {
		http.Error(resp, err.Error(), http.StatusBadRequest)
		return
	}

	resp.Header().Set("Content-Type", "application/json")
	resp.WriteHeader(http.StatusOK)
	err = json.NewEncoder(resp).Encode(customerGet)
	if err != nil {
		log.Println(err)
	}
}

//Update Customer
func (server *Server) UpdateCustomer(resp http.ResponseWriter, request *http.Request) {

	vars := mux.Vars(request)
	C_id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		http.Error(resp, err.Error(), http.StatusBadRequest)
		return
	}

	body, err := io.ReadAll(request.Body)
	if err != nil {
		http.Error(resp, err.Error(), http.StatusBadRequest)
		return
	}

	customer := entity.Customer{}
	err = json.Unmarshal(body, &customer)
	if err != nil {
		http.Error(resp, err.Error(), http.StatusBadRequest)
		return
	}
	updatedCustomer, err := customer.UpdateCustomer(server.DB, C_id)
	if err != nil {
		http.Error(resp, err.Error(), http.StatusBadRequest)
		return
	}

	resp.Header().Set("Content-Type", "application/json")
	resp.WriteHeader(http.StatusOK)
	err = json.NewEncoder(resp).Encode(updatedCustomer)
	if err != nil {
		log.Println(err)
	}
}

//Delet Customer
func (server *Server) DeleteCustomer(resp http.ResponseWriter, request *http.Request) {

	vars := mux.Vars(request)
	C_id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		http.Error(resp, err.Error(), http.StatusBadRequest)
		return
	}

	customer := entity.Customer{}
	_, err = customer.DeleteCustomer(server.DB, C_id)
	if err != nil {
		http.Error(resp, err.Error(), http.StatusBadRequest)
		return
	}

	resp.Header().Set("Content-Type", "application/json")
	resp.WriteHeader(http.StatusOK)
	err = json.NewEncoder(resp).Encode("")
	if err != nil {
		log.Println(err)
	}
}
