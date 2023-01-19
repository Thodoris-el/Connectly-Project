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

/*
This is The Conversation Controller.
Every Customer that has sent us a message, has a conversation.
The Conversation Stage guides us to the message that the chatbot will send
if "None" -> we hear only the trigger word "Buy"
if "Buy" -> prev msg contains the word "Buy"
if "Review" -> we get "Yes" in prev message and we have sent a revies
*/

//Create Conversation
func (server *Server) CreateConversation(resp http.ResponseWriter, request *http.Request) {

	body, err := io.ReadAll(request.Body)

	if err != nil {
		http.Error(resp, err.Error(), http.StatusBadRequest)
		return
	}

	conversation := entity.Conversation{}
	err = json.Unmarshal(body, &conversation)
	if err != nil {
		http.Error(resp, err.Error(), http.StatusBadRequest)
		return
	}
	conversation.CreatedAt = time.Now()
	conversation.UpdatedAt = time.Now()

	conversationCreated, err := conversation.SaveConversation(server.DB)

	if err != nil {
		http.Error(resp, err.Error(), http.StatusBadRequest)
		return
	}

	resp.Header().Set("Content-Type", "application/json")
	resp.WriteHeader(http.StatusOK)
	err = json.NewEncoder(resp).Encode(conversationCreated)
	if err != nil {
		log.Println(err)
	}
}

//Get All Conversations from DB
func (server *Server) GetConversation(resp http.ResponseWriter, request *http.Request) {

	conversation := entity.Conversation{}

	conversations, err := conversation.FindAllConversations(server.DB)
	if err != nil {
		http.Error(resp, err.Error(), http.StatusBadRequest)
		return
	}

	resp.Header().Set("Content-Type", "application/json")
	resp.WriteHeader(http.StatusOK)
	err = json.NewEncoder(resp).Encode(conversations)
	if err != nil {
		log.Println(err)
	}
}

//Get Conversation By ID
func (server *Server) GetConversationById(resp http.ResponseWriter, request *http.Request) {

	vars := mux.Vars(request)
	R_id, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		http.Error(resp, err.Error(), http.StatusBadRequest)
		return
	}

	conversation := entity.Conversation{}

	conversationGet, err := conversation.FindById(server.DB, int64(R_id))
	if err != nil {
		http.Error(resp, err.Error(), http.StatusBadRequest)
		return
	}

	resp.Header().Set("Content-Type", "application/json")
	resp.WriteHeader(http.StatusOK)
	err = json.NewEncoder(resp).Encode(conversationGet)
	if err != nil {
		log.Println(err)
	}
}

//Get Conversation by CustomerID
func (server *Server) GetConversationByCustomerId(resp http.ResponseWriter, request *http.Request) {

	vars := mux.Vars(request)
	C_id := vars["customer_id"]

	conversation := entity.Conversation{}

	conversationGet, err := conversation.FindByCustomerId(server.DB, C_id)
	if err != nil {
		http.Error(resp, err.Error(), http.StatusBadRequest)
		return
	}

	resp.Header().Set("Content-Type", "application/json")
	resp.WriteHeader(http.StatusOK)
	err = json.NewEncoder(resp).Encode(conversationGet)
	if err != nil {
		log.Println(err)
	}
}

//Update Conversation
func (server *Server) UpdateConversation(resp http.ResponseWriter, request *http.Request) {

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

	conversation := entity.Conversation{}
	err = json.Unmarshal(body, &conversation)
	if err != nil {
		http.Error(resp, err.Error(), http.StatusBadRequest)
		return
	}
	updatedConversation, err := conversation.UpdateConversation(server.DB, C_id)
	if err != nil {
		http.Error(resp, err.Error(), http.StatusBadRequest)
		return
	}

	resp.Header().Set("Content-Type", "application/json")
	resp.WriteHeader(http.StatusOK)
	err = json.NewEncoder(resp).Encode(updatedConversation)
	if err != nil {
		log.Println(err)
	}
}

//Delete Conversation
func (server *Server) DeleteConversation(resp http.ResponseWriter, request *http.Request) {

	vars := mux.Vars(request)
	C_id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		http.Error(resp, err.Error(), http.StatusBadRequest)
		return
	}

	conversation := entity.Conversation{}
	_, err = conversation.DeleteConversation(server.DB, C_id)
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
