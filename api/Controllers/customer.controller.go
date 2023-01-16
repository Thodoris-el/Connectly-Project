package controllers

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"

	entity "github.com/Thodoris-el/Coonectly-Project/api/Entity"
)

func (server *Server) CreateCustomer(resp http.ResponseWriter, request *http.Request) {

	body, err := io.ReadAll(request.Body)

	if err != nil {
		log.Println(err)
		resp.WriteHeader(400)
		resp.Write([]byte("Error while creating customer"))
		return
	}

	customer := entity.Customer{}
	err = json.Unmarshal(body, &customer)
	if err != nil {
		log.Println(err)
		resp.WriteHeader(400)
		resp.Write([]byte("Error while unmarshall"))
	}
	customer.CreatedAt = time.Now()
	customer.UpdatedAt = time.Now()

	customerCreated, err := customer.SaveCustomer(server.DB)

	if err != nil {
		log.Println(err)
		resp.WriteHeader(400)
		resp.Write([]byte("Error while creating customer"))
		return
	}
	resp.Header().Set("Content-Type", "application/json")
	resp.WriteHeader(http.StatusOK)
	err = json.NewEncoder(resp).Encode(customerCreated)
	if err != nil {
		log.Println(err)
	}

}

func (server *Server) GetCustomers(resp http.ResponseWriter, request *http.Request) {

	customer := entity.Customer{}

	customers, err := customer.FindAllCustomers(server.DB)
	if err != nil {
		log.Println(err)
		resp.WriteHeader(400)
		resp.Write([]byte("error while finding customers"))
	}

	resp.Header().Set("Content-Type", "application/json")
	resp.WriteHeader(http.StatusOK)
	err = json.NewEncoder(resp).Encode(customers)
	if err != nil {
		log.Println(err)
	}
}
