package controllers

func (server *Server) startRoutes() {
	//WebHook Verification
	server.Router.HandleFunc("/webhook", server.VerifyWebhook).Methods("GET")

	//Get Messasges From FB
	server.Router.HandleFunc("/webhook", server.HandleMessenger).Methods("POST")

	//Users endpoints
	server.Router.HandleFunc("/customer", server.CreateCustomer).Methods("POST")
	server.Router.HandleFunc("/customer", server.GetCustomers).Methods("GET")

	//Conversation Endpoints
	server.Router.HandleFunc("/conversation", server.CreateConversation).Methods("POST")
	server.Router.HandleFunc("/conversation", server.GetConversation).Methods("GET")
	server.Router.HandleFunc("/conversation/{id}", server.GetConversationById).Methods("GET")
	server.Router.HandleFunc("/conversation/customer/{customer_id}", server.GetConversationByCustomerId).Methods("GET")

	//Reviews endpoint
	server.Router.HandleFunc("/review", server.CreateReview).Methods("POST")
	server.Router.HandleFunc("/review", server.GetReviews).Methods("GET")
	server.Router.HandleFunc("/review/{id}", server.GetReviewById).Methods("GET")
	server.Router.HandleFunc("/review/customer/{customer_id}", server.GetReviewByCustomerId).Methods("GET")

}
