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

func TestCreateConversation(t *testing.T) {

	refreshTables()

	samples := []string{
		`{"facebook_id": "6706612322695175","stage": "buy","product": "bike"}`,
		`{"facebook_id": "8706612322695175","stage": "none","product": ""}`,
		`{"facebook_id": "9806612322695175","stage": "review","product": "bike"}`,
		`{"facebook_id": "6798612322695175","stage": "none","product": ""}`,
	}
	answers := []string{
		"buy",
		"none",
		"review",
		"none",
	}

	for i, tmp := range samples {
		req, err := http.NewRequest("POST", "/conversation", bytes.NewBufferString(tmp))
		if err != nil {
			t.Errorf("error: %v\n", err)
		}
		recorded := httptest.NewRecorder()
		handler := http.HandlerFunc(server.CreateConversation)
		handler.ServeHTTP(recorded, req)

		resp := make(map[string]interface{})
		err = json.Unmarshal(recorded.Body.Bytes(), &resp)
		if err != nil {
			t.Errorf("cant convert json: %v\n", err)
		}
		if resp["stage"] != answers[i] {
			t.Errorf("wrong values")
		}
	}
}

func TestGetConversation(t *testing.T) {

	refreshTables()

	_, err := createTwoConversations()
	if err != nil {
		t.Errorf("error creating  conversations")
	}

	req, err := http.NewRequest("GET", "/conversation", nil)
	if err != nil {
		t.Errorf("error: %v\n", err)
	}
	recorded := httptest.NewRecorder()
	handler := http.HandlerFunc(server.GetConversation)
	handler.ServeHTTP(recorded, req)

	var conversations []entity.Conversation
	err = json.Unmarshal(recorded.Body.Bytes(), &conversations)
	if err != nil {
		t.Errorf("error while unmarshal")
	}
	if len(conversations) != 2 {
		t.Errorf("wrong number of conversations")
	}
}

func TestGetConversationById(t *testing.T) {

	refreshTables()

	_, err := createAConversation()
	if err != nil {
		t.Errorf("error creating a conversation")
	}

	req, err := http.NewRequest("GET", "/conversation", nil)
	if err != nil {
		t.Errorf("error: %v\n", err)
	}
	req = mux.SetURLVars(req, map[string]string{"id": "1"})
	recorded := httptest.NewRecorder()
	handler := http.HandlerFunc(server.GetConversationById)
	handler.ServeHTTP(recorded, req)

	responseMap := make(map[string]interface{})
	err = json.Unmarshal(recorded.Body.Bytes(), &responseMap)
	if err != nil {
		t.Errorf("error while unmarshal")
	}
	if responseMap["stage"] != "None" {
		t.Errorf("wrong value")
	}
}

func TestGetConversationByCustomerId(t *testing.T) {

	refreshTables()

	_, err := createAConversation()
	if err != nil {
		t.Errorf("error creating a conversation")
	}

	req, err := http.NewRequest("GET", "/conversation/customer", nil)
	if err != nil {
		t.Errorf("error: %v\n", err)
	}
	req = mux.SetURLVars(req, map[string]string{"customer_id": "6706612322695175"})
	recorded := httptest.NewRecorder()
	handler := http.HandlerFunc(server.GetConversationByCustomerId)
	handler.ServeHTTP(recorded, req)

	responseMap := make(map[string]interface{})
	err = json.Unmarshal(recorded.Body.Bytes(), &responseMap)
	if err != nil {
		t.Errorf("error while unmarshal")
	}
	if responseMap["stage"] != "None" {
		t.Errorf("wrong value")
	}
}

func TestUpdateConversationC(t *testing.T) {

	refreshTables()

	_, err := createAConversation()
	if err != nil {
		t.Errorf("error creating a conversation")
	}
	sample := `{"stage": "Buy"}`
	req, err := http.NewRequest("PUT", "/conversation", bytes.NewBufferString(sample))
	if err != nil {
		t.Errorf("error: %v\n", err)
	}

	req = mux.SetURLVars(req, map[string]string{"id": "1"})
	recorded := httptest.NewRecorder()
	handler := http.HandlerFunc(server.UpdateConversation)
	handler.ServeHTTP(recorded, req)

	responseMap := make(map[string]interface{})
	err = json.Unmarshal(recorded.Body.Bytes(), &responseMap)
	if err != nil {
		t.Errorf("error while unmarshal")
	}
	if responseMap["stage"] != "Buy" {
		t.Errorf("wrong value")
	}
}

func TestDeleteConversationc(t *testing.T) {

	refreshTables()

	_, err := createAConversation()
	if err != nil {
		t.Errorf("error creating a conversation")
	}

	req, err := http.NewRequest("DELETE", "/conversation", nil)
	if err != nil {
		t.Errorf("error: %v\n", err)
	}
	req = mux.SetURLVars(req, map[string]string{"id": "1"})
	recorded := httptest.NewRecorder()
	handler := http.HandlerFunc(server.DeleteConversation)
	handler.ServeHTTP(recorded, req)

	var conversations entity.Conversation
	getC, err := conversations.FindAllConversations(server.DB)
	if err != nil {
		t.Errorf("error while getting conversations")
	}
	if len(*getC) != 0 {
		t.Errorf("wrong number of conversations")
	}
}
