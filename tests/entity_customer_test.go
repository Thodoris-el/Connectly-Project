package tests

import (
	"testing"
	"time"

	entity "github.com/Thodoris-el/Coonectly-Project/api/Entity"
)

func TestSaveCustomer(t *testing.T) {

	refreshTables()

	testCustomer := entity.Customer{
		First_name:  "John",
		Last_name:   "Dir",
		Facebook_id: "6706612322695175",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	savedCutomer, err := testCustomer.SaveCustomer(server.DB)

	if err != nil {
		t.Errorf("error saving the customer: %v\n", err)
		return
	}

	if !(testCustomer.First_name == savedCutomer.First_name && testCustomer.Last_name == savedCutomer.Last_name && testCustomer.Facebook_id == savedCutomer.Facebook_id) {
		t.Errorf("wrong values")
	}
}

func TestFindAllCustomers(t *testing.T) {

	refreshTables()

	_, err := createTwoCustomers()
	testCustomer := entity.Customer{}

	if err != nil {
		t.Errorf("error creating the customer: %v\n", err)
		return
	}

	getC, err := testCustomer.FindAllCustomers(server.DB)
	if err != nil {
		t.Errorf("error fetching the customer: %v\n", err)
		return
	}

	if len(*getC) != 2 {
		t.Errorf("wrong number of customers")
	}
}

func TestFindCustomerByID(t *testing.T) {

	refreshTables()

	testCustomer, err := createACustomer()

	if err != nil {
		t.Errorf("error creating the customer: %v\n", err)
		return
	}

	getC, err := testCustomer.FindCustomerByID(server.DB, 1)
	if err != nil {
		t.Errorf("error fetching the customer: %v\n", err)
		return
	}

	if !(testCustomer.First_name == getC.First_name && testCustomer.Last_name == getC.Last_name && testCustomer.Facebook_id == getC.Facebook_id) {
		t.Errorf("wrong values")
	}
}

func TestUpdateCustomer(t *testing.T) {

	refreshTables()

	testCustomer, err := createACustomer()

	if err != nil {
		t.Errorf("error creating the customer: %v\n", err)
		return
	}

	testCustomer.First_name = "Nick"
	getC, err := testCustomer.UpdateCustomer(server.DB, 1)
	if err != nil {
		t.Errorf("error fetching the customer: %v\n", err)
		return
	}

	if !(testCustomer.First_name == getC.First_name && testCustomer.Last_name == getC.Last_name && testCustomer.Facebook_id == getC.Facebook_id) {
		t.Errorf("wrong values")
	}
}

func TestDeleteCustomer(t *testing.T) {

	refreshTables()

	testCustomer, err := createACustomer()

	if err != nil {
		t.Errorf("error creating the customer: %v\n", err)
		return
	}

	getC, err := testCustomer.DeleteCustomer(server.DB, 1)
	if err != nil {
		t.Errorf("error fetching the customer: %v\n", err)
		return
	}

	if getC != 1 {
		t.Errorf("wrong values")
	}
}
