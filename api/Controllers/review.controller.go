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

func (server *Server) CreateReview(resp http.ResponseWriter, request *http.Request) {

	body, err := ioutil.ReadAll(request.Body)

	if err != nil {
		log.Println(err)
		resp.WriteHeader(400)
		resp.Write([]byte("Error while creating review"))
		return
	}

	review := entity.Review{}
	err = json.Unmarshal(body, &review)
	if err != nil {
		log.Println(err)
		resp.WriteHeader(400)
		resp.Write([]byte("Error while unmarshall"))
	}
	review.CreatedAt = time.Now()
	review.UpdatedAt = time.Now()

	reviewCreated, err := review.SaveReview(server.DB)

	if err != nil {
		log.Println(err)
		resp.WriteHeader(400)
		resp.Write([]byte("Error while creating review"))
		return
	}

	resp.WriteHeader(200)
	err = json.NewEncoder(resp).Encode(reviewCreated)
	if err != nil {
		log.Println(err)
	}
}

func (server *Server) GetReviews(resp http.ResponseWriter, request *http.Request) {

	review := entity.Review{}

	reviews, err := review.FindAllReviews(server.DB)
	if err != nil {
		log.Println(err)
		resp.WriteHeader(400)
		resp.Write([]byte("error while finding reviews"))
	}

	resp.WriteHeader(200)
	err = json.NewEncoder(resp).Encode(reviews)
	if err != nil {
		log.Println(err)
	}
}

func (server *Server) GetPostById(resp http.ResponseWriter, request *http.Request) {

	vars := mux.Vars(request)
	R_id, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		log.Println("bad request")
		resp.WriteHeader(http.StatusBadRequest)
		resp.Write([]byte("bad request"))
		return
	}

	review := entity.Review{}

	reviewGet, err := review.FindById(server.DB, int64(R_id))
	if err != nil {
		log.Println(err)
		resp.WriteHeader(400)
		resp.Write([]byte("internal server error"))
		return
	}

	resp.WriteHeader(200)
	err = json.NewEncoder(resp).Encode(reviewGet)
	if err != nil {
		log.Println(err)
	}
}

func (server *Server) GetPostByCustomerId(resp http.ResponseWriter, request *http.Request) {

	vars := mux.Vars(request)
	C_id := vars["customer_id"]

	review := entity.Review{}

	reviewGet, err := review.FindByCustomerId(server.DB, C_id)
	if err != nil {
		log.Println(err)
		resp.WriteHeader(400)
		resp.Write([]byte("internal server error"))
		return
	}

	resp.WriteHeader(200)
	err = json.NewEncoder(resp).Encode(reviewGet)
	if err != nil {
		log.Println(err)
	}
}

func (server *Server) AddReview(senderID, text, score string) error {

	var err error
	new_review := entity.Review{}

	//Add values to fields
	new_review.Customer_id = senderID
	new_review.Text = text
	new_review.Score, err = strconv.Atoi(score)
	if err != nil {
		log.Println("error converting score to integer")
		return err
	}
	new_review.CreatedAt = time.Now()
	new_review.UpdatedAt = time.Now()

	if text == "" {
		switch score {
		case "1":
			new_review.Text = "Very Dissatisfied"
		case "2":
			new_review.Text = "Dissatisfied"
		case "3":
			new_review.Text = "neutral"
		case "4":
			new_review.Text = "Satisfied"
		default:
			new_review.Text = "Very Satisfied"
		}
	}

	_, err = new_review.SaveReview(server.DB)
	if err != nil {
		log.Println("error while savung review")
		return err
	}
	return nil
}
