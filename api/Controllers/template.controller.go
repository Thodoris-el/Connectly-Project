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
This is the Template Controller
Implements:
-> Create a template
-> Get All templates
-> Get customer by ID
-> Get template by language
-> Update customer
-> Delete customer
*/

//Create Template
func (server *Server) CreateTemplate(resp http.ResponseWriter, request *http.Request) {

	body, err := io.ReadAll(request.Body)

	if err != nil {
		http.Error(resp, err.Error(), http.StatusBadRequest)
		return
	}

	template := entity.Template{}
	err = json.Unmarshal(body, &template)
	if err != nil {
		http.Error(resp, err.Error(), http.StatusBadRequest)
		return
	}
	template.CreatedAt = time.Now()
	template.UpdatedAt = time.Now()

	templateCreated, err := template.SaveTemplate(server.DB)

	if err != nil {
		http.Error(resp, err.Error(), http.StatusBadRequest)
		return
	}

	resp.Header().Set("Content-Type", "application/json")
	resp.WriteHeader(http.StatusOK)
	err = json.NewEncoder(resp).Encode(templateCreated)
	if err != nil {
		log.Println(err)
	}
}

//Get All Templates from DB
func (server *Server) GetTemplate(resp http.ResponseWriter, request *http.Request) {

	template := entity.Template{}

	templates, err := template.FindAllTemplates(server.DB)
	if err != nil {
		http.Error(resp, err.Error(), http.StatusBadRequest)
		return
	}

	resp.Header().Set("Content-Type", "application/json")
	resp.WriteHeader(http.StatusOK)
	err = json.NewEncoder(resp).Encode(templates)
	if err != nil {
		log.Println(err)
	}
}

//Get Template By ID from DB
func (server *Server) GetTemplateById(resp http.ResponseWriter, request *http.Request) {

	vars := mux.Vars(request)
	T_id, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		http.Error(resp, err.Error(), http.StatusBadRequest)
		return
	}

	template := entity.Template{}

	templateGet, err := template.FindById(server.DB, int64(T_id))
	if err != nil {
		http.Error(resp, err.Error(), http.StatusBadRequest)
		return
	}

	resp.Header().Set("Content-Type", "application/json")
	resp.WriteHeader(http.StatusOK)
	err = json.NewEncoder(resp).Encode(templateGet)
	if err != nil {
		log.Println(err)
	}
}

//Get Template by customerID from DB
func (server *Server) GetTemplateByLanguage(resp http.ResponseWriter, request *http.Request) {

	vars := mux.Vars(request)
	lang := vars["language"]

	template := entity.Template{}

	templateGet, err := template.FindByLanguage(server.DB, lang)
	if err != nil {
		http.Error(resp, err.Error(), http.StatusBadRequest)
		return
	}

	resp.Header().Set("Content-Type", "application/json")
	resp.WriteHeader(http.StatusOK)
	err = json.NewEncoder(resp).Encode(templateGet)
	if err != nil {
		log.Println(err)
	}
}

//Update Template
func (server *Server) UpdateTemplate(resp http.ResponseWriter, request *http.Request) {

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

	template := entity.Template{}
	err = json.Unmarshal(body, &template)
	if err != nil {
		http.Error(resp, err.Error(), http.StatusBadRequest)
		return
	}
	updatedTemplate, err := template.UpdateTemplate(server.DB, C_id)
	if err != nil {
		http.Error(resp, err.Error(), http.StatusBadRequest)
		return
	}

	resp.Header().Set("Content-Type", "application/json")
	resp.WriteHeader(http.StatusOK)
	err = json.NewEncoder(resp).Encode(updatedTemplate)
	if err != nil {
		log.Println(err)
	}
}

//Delete Template
func (server *Server) DeleteTemplate(resp http.ResponseWriter, request *http.Request) {

	vars := mux.Vars(request)
	C_id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		http.Error(resp, err.Error(), http.StatusBadRequest)
		return
	}

	template := entity.Template{}
	_, err = template.DeleteTemplate(server.DB, C_id)
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
