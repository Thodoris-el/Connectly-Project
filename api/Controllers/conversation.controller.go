package controllers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"

	entity "github.com/Thodoris-el/Coonectly-Project/api/Entity"
	"github.com/gorilla/mux"
)

func (server *Server) CreateConversation(resp http.ResponseWriter, request *http.Request) {

	body, err := ioutil.ReadAll(request.Body)

	if err != nil {
		log.Println(err)
		resp.WriteHeader(400)
		resp.Write([]byte("Error while creating conversation"))
		return
	}

	conversation := entity.Conversation{}
	err = json.Unmarshal(body, &conversation)
	if err != nil {
		log.Println(err)
		resp.WriteHeader(400)
		resp.Write([]byte("Error while unmarshall"))
	}
	conversation.CreatedAt = time.Now()
	conversation.UpdatedAt = time.Now()

	conversationCreated, err := conversation.SaveConversation(server.DB)

	if err != nil {
		log.Println(err)
		resp.WriteHeader(400)
		resp.Write([]byte("Error while creating conversation"))
		return
	}

	resp.WriteHeader(200)
	err = json.NewEncoder(resp).Encode(conversationCreated)
	if err != nil {
		log.Println(err)
	}
}

func (server *Server) GetConversation(resp http.ResponseWriter, request *http.Request) {

	review := entity.Conversation{}

	conversations, err := review.FindAllConversations(server.DB)
	if err != nil {
		log.Println(err)
		resp.WriteHeader(400)
		resp.Write([]byte("error while finding conversations"))
	}

	resp.WriteHeader(200)
	err = json.NewEncoder(resp).Encode(conversations)
	if err != nil {
		log.Println(err)
	}
}

func (server *Server) GetConversationById(resp http.ResponseWriter, request *http.Request) {

	vars := mux.Vars(request)
	R_id, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		log.Println("bad request")
		resp.WriteHeader(http.StatusBadRequest)
		resp.Write([]byte("bad request"))
		return
	}

	conversation := entity.Conversation{}

	conversationGet, err := conversation.FindById(server.DB, int64(R_id))
	if err != nil {
		log.Println(err)
		resp.WriteHeader(400)
		resp.Write([]byte("internal server error"))
		return
	}

	resp.WriteHeader(200)
	err = json.NewEncoder(resp).Encode(conversationGet)
	if err != nil {
		log.Println(err)
	}
}

func (server *Server) GetConversationByCustomerId(resp http.ResponseWriter, request *http.Request) {

	vars := mux.Vars(request)
	C_id := vars["customer_id"]

	conversation := entity.Conversation{}

	conversationGet, err := conversation.FindByCustomerId(server.DB, C_id)
	if err != nil {
		log.Println(err)
		resp.WriteHeader(400)
		resp.Write([]byte("internal server error"))
		return
	}

	resp.WriteHeader(200)
	err = json.NewEncoder(resp).Encode(conversationGet)
	if err != nil {
		log.Println(err)
	}
}
