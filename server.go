package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func HandleMessenger(resp http.ResponseWriter, request *http.Request) {
	secretKey := "secret_token123"
	if request.Method == "GET" {
		mode := request.URL.Query().Get("hub.mode")
		challenge := request.URL.Query().Get("hub.challenge")
		token := request.URL.Query().Get("hub.verify_token")
		fmt.Println(request.Body)
		fmt.Println(mode)
		fmt.Println(challenge)
		fmt.Println(token)
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
	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		log.Printf("failed to read body: %v", err)
	}

	// print the body
	log.Printf(string(body))

}

// Initialize request

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/webhook", HandleMessenger).Methods("POST", "GET")
	router.HandleFunc("/", HandleMessenger).Methods("POST", "GET")
	port := ":8000"
	log.Printf("Server started on %s", port)
	log.Fatal(http.ListenAndServe(port, router))
}
