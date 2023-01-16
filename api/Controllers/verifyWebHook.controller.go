package controllers

import (
	"fmt"
	"net/http"
)

func (server *Server) VerifyWebhook(resp http.ResponseWriter, request *http.Request) {
	secretKey := "secret_token123"
	if request.Method == "GET" {
		mode := request.URL.Query().Get("hub.mode")
		challenge := request.URL.Query().Get("hub.challenge")
		token := request.URL.Query().Get("hub.verify_token")
		if mode != "" && token != "" {
			if token == secretKey && mode == "subscribe" {
				fmt.Println("WEBHOOK_VERIFIED")
				resp.WriteHeader(200)
				resp.Write([]byte(challenge))
				return
			}
		}
		resp.WriteHeader(400)
		resp.Write([]byte(`Bad token`))
	}
}
