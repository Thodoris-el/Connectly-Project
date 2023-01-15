package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	entity "github.com/Thodoris-el/Coonectly-Project/api/Entity"
	"github.com/gorilla/mux"
)

func VerifyWebhook(resp http.ResponseWriter, request *http.Request) {
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
}

func HandleMessenger(resp http.ResponseWriter, request *http.Request) {
	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		log.Printf("failed to read body: %v", err)
	}

	var facebookPost entity.FacebookMessage
	err = json.Unmarshal(body, &facebookPost)
	// print the body
	if err != nil {
		log.Println(err)
		resp.WriteHeader(400)
		resp.Write([]byte("Error while unmarshal request"))
		return
	}
	if facebookPost.Object != "page" {
		log.Panicln("No fb object page")
		resp.WriteHeader(400)
		resp.Write([]byte("No fb object page"))
	}
	fbEntry := facebookPost.Entry
	var fbMess []entity.MessagingType
	for _, everyEntry := range fbEntry {
		fbMess = everyEntry.Messaging
		var sender entity.SenderType
		var recipient entity.RecipientType
		var message entity.MessageType
		for _, everyfbMess := range fbMess {
			sender = everyfbMess.Sender
			recipient = everyfbMess.Recipient
			message = everyfbMess.Message
			if message.Attachments == nil {
				handleMessage(sender.ID, "hello back")
				log.Println(sender, recipient, message.Text)
			} else {
				attachment := message.Attachments
				for _, everyAttachment := range attachment {
					switch everyAttachment.Type {
					case "template":
						log.Println(sender, recipient, everyAttachment.Payload.Product.Title)
					default:
						log.Println(sender, recipient, "image: ", everyAttachment.Payload.URL)
					}
				}
			}
		}
	}
}

// Handles messages
func handleMessage(senderId, message string) error {
	if len(message) == 0 {
		log.Println("No message found.")
	}
	response := entity.MessangeSent{
		Recipient: entity.RecipientType{
			ID: senderId,
		},
		Message: entity.ResMessageType{
			Text: message,
		},
	}
	data, err := json.Marshal(response)

	if err != nil {
		log.Println("Marshal error: %s", err)
		return err
	}
	url_response := "https://graph.facebook.com/v2.6/me/messages"
	FACEBOOK_TOKEN := "EAAMdOMePqfEBAN9hZAQStzAHqpF3tY54rRzYJwqZBcajQycusrNN6OxYw6dcxSa5ZAqbkZBfDrmaiwYFQ49jeq8SoPmoglcxy6pMce5y7H7Qc4l25bX1KkgK2zMSkK3mnyMjbTWDIYG1ZCxLZAQFangm93cFZB2UavrrBhDFR2TVO01bmS0hpUy6qmiQaqShwgZD"
	url_response = fmt.Sprintf("%s?access_token=%s", url_response, FACEBOOK_TOKEN)
	log.Println("URL: %s", url_response)
	req, err := http.NewRequest("POST", url_response, bytes.NewBuffer(data))
	if err != nil {
		log.Println("Failed making request: %s", err)
		return err
	}
	req.Header.Add("Content-Type", "application/json")
	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		log.Printf("Failed doing request: %s", err)
		return err
	}
	log.Printf("MESSAGE SENT?\n%#v", res)
	return nil
}

// Initialize request

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/webhook", VerifyWebhook).Methods("GET")
	router.HandleFunc("/webhook", HandleMessenger).Methods("POST")
	port := ":8000"
	log.Printf("Server started on %s", port)
	log.Fatal(http.ListenAndServe(port, router))
}
