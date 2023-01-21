package controllers

import (
	"net/http"
)

func (server *Server) startRoutes() {

	//home route -> just responds that the server is working -> used for healthcheck
	server.Router.HandleFunc("/", func(resp http.ResponseWriter, request *http.Request) {
		resp.WriteHeader(http.StatusOK)
		resp.Write([]byte("server working"))
	}).Methods("GET")

	//Facebook WebHook Verification
	server.Router.HandleFunc("/facebook/webhook", server.VerifyWebhook).Methods("GET")
	//Get Messasges From FB
	server.Router.HandleFunc("/facebook/webhook", server.HandleReceivedFacebookMessage).Methods("POST")

	//Users endpoints
	server.Router.HandleFunc("/customer", server.CreateCustomer).Methods("POST")
	server.Router.HandleFunc("/customer", server.GetCustomers).Methods("GET")
	server.Router.HandleFunc("/customer/{id}", server.GetCustomerById).Methods("GET")
	server.Router.HandleFunc("/customer/facebook/{facebook_id}", server.GetCustomerByFacebookId).Methods("GET")
	server.Router.HandleFunc("/customer/{id}", server.UpdateCustomer).Methods("PUT")
	server.Router.HandleFunc("/customer/{id}", server.DeleteCustomer).Methods("DELETE")

	//Conversation Endpoints
	server.Router.HandleFunc("/conversation", server.CreateConversation).Methods("POST")
	server.Router.HandleFunc("/conversation", server.GetConversation).Methods("GET")
	server.Router.HandleFunc("/conversation/{id}", server.GetConversationById).Methods("GET")
	server.Router.HandleFunc("/conversation/customer/{customer_id}", server.GetConversationByCustomerId).Methods("GET")
	server.Router.HandleFunc("/conversation/{id}", server.UpdateConversation).Methods("PUT")
	server.Router.HandleFunc("/conversation/{id}", server.DeleteConversation).Methods("DELETE")

	//Reviews endpoints
	server.Router.HandleFunc("/review", server.CreateReview).Methods("POST")
	server.Router.HandleFunc("/review", server.GetReviews).Methods("GET")
	server.Router.HandleFunc("/review/{id}", server.GetReviewById).Methods("GET")
	server.Router.HandleFunc("/review/customer/{customer_id}", server.GetReviewByCustomerId).Methods("GET")
	server.Router.HandleFunc("/review/product/{product}", server.GetReviewByProduct).Methods("GET")
	server.Router.HandleFunc("/review/{id}", server.UpdateReview).Methods("PUT")
	server.Router.HandleFunc("/review/{id}", server.DeleteReview).Methods("DELETE")

	//Template endpoints
	server.Router.HandleFunc("/template", server.CreateTemplate).Methods("POST")
	server.Router.HandleFunc("/template", server.GetTemplate).Methods("GET")
	server.Router.HandleFunc("/template/{id}", server.GetTemplateById).Methods("GET")
	server.Router.HandleFunc("/template/language/{language}", server.GetTemplateByLanguage).Methods("GET")
	server.Router.HandleFunc("/template/{id}", server.UpdateTemplate).Methods("PUT")
	server.Router.HandleFunc("/template/{id}", server.DeleteTemplate).Methods("DELETE")
}
