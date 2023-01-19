package tests

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	entity "github.com/Thodoris-el/Coonectly-Project/api/Entity"
	"github.com/gorilla/mux"
)

func TestCreateReview(t *testing.T) {

	refreshTables()

	samples := []string{
		`{"customer_id": "6706612322695175","text": "buy","score": 1}`,
		`{"customer_id": "8706612322695175","text": "none","score": 2}`,
		`{"customer_id": "9806612322695175","text": "review","score": 3}`,
		`{"customer_id": "6798612322695175","text": "none","score": 4}`,
	}
	answers := []string{
		"buy",
		"none",
		"review",
		"none",
	}

	for i, tmp := range samples {
		req, err := http.NewRequest("POST", "/review", bytes.NewBufferString(tmp))
		if err != nil {
			t.Errorf("error: %v\n", err)
		}
		recorded := httptest.NewRecorder()
		handler := http.HandlerFunc(server.CreateReview)
		handler.ServeHTTP(recorded, req)

		resp := make(map[string]interface{})
		err = json.Unmarshal(recorded.Body.Bytes(), &resp)
		if err != nil {
			t.Errorf("cant convert json: %v\n", err)
		}
		if resp["text"] != answers[i] {
			t.Errorf("wrong values")
		}
	}
}

func TestGetReviews(t *testing.T) {

	refreshTables()

	_, err := createTwoReviews()
	if err != nil {
		t.Errorf("error creating  reviews")
	}

	req, err := http.NewRequest("GET", "/review", nil)
	if err != nil {
		t.Errorf("error: %v\n", err)
	}
	recorded := httptest.NewRecorder()
	handler := http.HandlerFunc(server.GetReviews)
	handler.ServeHTTP(recorded, req)

	var reviews []entity.Review
	err = json.Unmarshal(recorded.Body.Bytes(), &reviews)
	if err != nil {
		t.Errorf("error while unmarshal")
	}
	if len(reviews) != 2 {
		t.Errorf("wrong number of reviews")
	}
}

func TestGetReviewById(t *testing.T) {

	refreshTables()

	_, err := createAReview()
	if err != nil {
		t.Errorf("error creating a review")
	}

	req, err := http.NewRequest("GET", "/review", nil)
	if err != nil {
		t.Errorf("error: %v\n", err)
	}
	req = mux.SetURLVars(req, map[string]string{"id": "1"})
	recorded := httptest.NewRecorder()
	handler := http.HandlerFunc(server.GetReviewById)
	handler.ServeHTTP(recorded, req)

	responseMap := make(map[string]interface{})
	err = json.Unmarshal(recorded.Body.Bytes(), &responseMap)
	if err != nil {
		t.Errorf("error while unmarshal")
	}
	if responseMap["text"] != "1" {
		t.Errorf("wrong value")
	}
}

func TestGetReviewByCustomerId(t *testing.T) {

	refreshTables()

	_, err := createAReview()
	if err != nil {
		t.Errorf("error creating a review")
	}

	req, err := http.NewRequest("GET", "/review/customer", nil)
	if err != nil {
		t.Errorf("error: %v\n", err)
	}
	req = mux.SetURLVars(req, map[string]string{"customer_id": "6706612322695175"})
	recorded := httptest.NewRecorder()
	handler := http.HandlerFunc(server.GetReviewByCustomerId)
	handler.ServeHTTP(recorded, req)

	var reviews []entity.Review
	err = json.Unmarshal(recorded.Body.Bytes(), &reviews)
	if err != nil {
		log.Println(err)
		t.Errorf("error while unmarshal")
	}
	if len(reviews) != 1 {
		t.Errorf("wrong value")
	}
}

func TestGetReviewByProduct(t *testing.T) {

	refreshTables()

	_, err := createAReview()
	if err != nil {
		t.Errorf("error creating a review")
	}

	req, err := http.NewRequest("GET", "/review/product", nil)
	if err != nil {
		t.Errorf("error: %v\n", err)
	}
	req = mux.SetURLVars(req, map[string]string{"product": "car"})
	recorded := httptest.NewRecorder()
	handler := http.HandlerFunc(server.GetReviewByProduct)
	handler.ServeHTTP(recorded, req)

	var reviews []entity.Review
	err = json.Unmarshal(recorded.Body.Bytes(), &reviews)
	if err != nil {
		log.Println(err)
		t.Errorf("error while unmarshal")
	}
	if len(reviews) != 1 {
		t.Errorf("wrong value")
	}
}

func TestUpdateReviewC(t *testing.T) {

	refreshTables()

	_, err := createAReview()
	if err != nil {
		t.Errorf("error creating a review")
	}
	sample := `{"text": "Buy"}`
	req, err := http.NewRequest("PUT", "/review", bytes.NewBufferString(sample))
	if err != nil {
		t.Errorf("error: %v\n", err)
	}

	req = mux.SetURLVars(req, map[string]string{"id": "1"})
	recorded := httptest.NewRecorder()
	handler := http.HandlerFunc(server.UpdateReview)
	handler.ServeHTTP(recorded, req)

	responseMap := make(map[string]interface{})
	err = json.Unmarshal(recorded.Body.Bytes(), &responseMap)
	if err != nil {
		t.Errorf("error while unmarshal")
	}
	if responseMap["text"] != "Buy" {
		t.Errorf("wrong value")
	}
}

func TestDeleteReviewc(t *testing.T) {

	refreshTables()

	_, err := createAReview()
	if err != nil {
		t.Errorf("error creating a review")
	}

	req, err := http.NewRequest("DELETE", "/review", nil)
	if err != nil {
		t.Errorf("error: %v\n", err)
	}
	req = mux.SetURLVars(req, map[string]string{"id": "1"})
	recorded := httptest.NewRecorder()
	handler := http.HandlerFunc(server.DeleteReview)
	handler.ServeHTTP(recorded, req)

	var reviews entity.Review
	getR, err := reviews.FindAllReviews(server.DB)
	if err != nil {
		t.Errorf("error while getting reviews")
	}
	if len(*getR) != 0 {
		t.Errorf("wrong number of reviews")
	}
}
