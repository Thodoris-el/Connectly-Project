package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
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
		err = json.Unmarshal([]byte(recorded.Body.String()), &resp)
		if err != nil {
			t.Errorf("cant vonvert json: %v\n", err)
		}
		if resp["first_name"] != answers[i] {
			t.Errorf("wrong values")
		}
	}
}
