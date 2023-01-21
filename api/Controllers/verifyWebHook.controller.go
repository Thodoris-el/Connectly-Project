package controllers

import (
	"log"
	"net/http"
	"os"
)

//handler for webhook verification
func (server *Server) VerifyWebhook(resp http.ResponseWriter, request *http.Request) {
	//get secret token from env
	secretKey := os.Getenv("SECRET_TOKEN")
	mode := request.URL.Query().Get("hub.mode")
	challenge := request.URL.Query().Get("hub.challenge")
	token := request.URL.Query().Get("hub.verify_token")
	if mode != "" && token != "" {
		if token == secretKey && mode == "subscribe" {
			log.Println("WEBHOOK_VERIFIED")
			resp.WriteHeader(200)
			resp.Write([]byte(challenge))
			return
		}
	}
	resp.WriteHeader(400)
	resp.Write([]byte(`Bad token`))
}
