package controllers

func (server *Server) startRoutes() {
	server.Router.HandleFunc("/webhook", VerifyWebhook).Methods("GET")
	server.Router.HandleFunc("/webhook", HandleMessenger).Methods("POST")

	server.Router.HandleFunc("/users", server.CreateCustomer).Methods("POST")
	server.Router.HandleFunc("/users", server.GetCustomers).Methods("GET")
}
