package controllers

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"

	entity "github.com/Thodoris-el/Coonectly-Project/api/Entity"
)

// Handles messages
func handleMessageWithQuickReply(senderId, message string) error {
	if len(message) == 0 {
		log.Println("No message found.")
		err := errors.New("no message found")
		return err
	}

	var quickReply []entity.QuickReplyType

	var quickReplytmp entity.QuickReplyType
	quickReplytmp.Content_Type = "text"
	quickReplytmp.Title = "Yes"
	quickReplytmp.Payload = "<DEVELOPER_DEFINED_PAYLOAD>"
	quickReplytmp.Image = "http://example.com/img/red.png"

	quickReply = append(quickReply, quickReplytmp)

	quickReplytmp.Content_Type = "text"
	quickReplytmp.Title = "No"
	quickReplytmp.Payload = "<DEVELOPER_DEFINED_PAYLOAD>"
	quickReplytmp.Image = "http://example.com/img/green.png"

	quickReply = append(quickReply, quickReplytmp)

	response := entity.MessangeSent{
		Recipient: entity.RecipientType{
			ID: senderId,
		},
		Message: entity.ResMessageType{
			Text:          message,
			Quick_Replies: quickReply,
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

// Handles messages
func handleMessageWithoutQuickReply(senderId, message string) error {
	if len(message) == 0 {
		log.Println("No message found.")
		err := errors.New("no message found")
		return err
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
