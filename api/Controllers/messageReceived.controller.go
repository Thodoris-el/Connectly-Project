package controllers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	entity "github.com/Thodoris-el/Coonectly-Project/api/Entity"
)

func HandleMessenger(resp http.ResponseWriter, request *http.Request) {
	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		log.Printf("failed to read body: %v", err)
	}

	var facebookPost entity.FacebookMessage
	err = json.Unmarshal(body, &facebookPost)
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
