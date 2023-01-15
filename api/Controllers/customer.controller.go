package controllers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	entity "github.com/Thodoris-el/Coonectly-Project/api/Entity"
)

func (server *Server) CreateCustomer(resp http.ResponseWriter, request *http.Request) {

	body, err := ioutil.ReadAll(request.Body)

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
	resp.WriteHeader(200)
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

	resp.WriteHeader(200)
	err = json.NewEncoder(resp).Encode(customers)
	if err != nil {
		log.Println(err)
	}
}
