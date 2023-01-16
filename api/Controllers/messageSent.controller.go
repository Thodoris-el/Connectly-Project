package controllers

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"

	entity "github.com/Thodoris-el/Coonectly-Project/api/Entity"
)

// Handles messages
func handleMessageWithQuickReply(senderId, message string) error {
	if len(message) == 0 {
		log.Println("No message found.")
		err := errors.New("no message found")
		return err
	}

	//Create Quick reply for product buy verification
	var quickReply []entity.QuickReplyType

	var quickReplytmp entity.QuickReplyType
	quickReplytmp.Content_Type = "text"
	quickReplytmp.Title = "Yes"
	quickReplytmp.Payload = "Buy Product"
	quickReplytmp.Image = "http://example.com/img/red.png"

	quickReply = append(quickReply, quickReplytmp)

	quickReplytmp.Content_Type = "text"
	quickReplytmp.Title = "No"
	quickReplytmp.Payload = "Don't Buy Product"
	quickReplytmp.Image = "http://example.com/img/green.png"

	quickReply = append(quickReply, quickReplytmp)

	//Create a Response to send to FB
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
	fb_token := os.Getenv("FACEBOOK_TOKEN")
	url_response = fmt.Sprintf("%s?access_token=%s", url_response, fb_token)
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
	fb_token := os.Getenv("FACEBOOK_TOKEN")
	url_response = fmt.Sprintf("%s?access_token=%s", url_response, fb_token)
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

func SendReviewTemplate(senderId string) error {

	followUp := entity.FollowUpType{
		Type:        "free_form",
		Placeholder: "Give additional feedback",
	}

	question := entity.QuestionType{
		ID:           "myquestion1",
		Type:         "csat",
		Title:        "How would you rate our product?",
		Score_Label:  "neg_pos",
		Score_Option: "five_stars",
		FollowUp:     followUp,
	}
	questions := []entity.QuestionType{}
	questions = append(questions, question)

	feedback := entity.FeedbackType{
		Questions: questions,
	}
	feedbacks := []entity.FeedbackType{}
	feedbacks = append(feedbacks, feedback)

	business := entity.BusinessPrivacyType{
		Url: "https://www.example.com",
	}
	payload := entity.PayloadTypeSendTemplate{
		Template_Type:    "customer_feedback",
		Title:            "Review Product",
		Subtitle:         "Let us know about the product",
		Button_Title:     "Review Product",
		Feedbavk_Screens: feedbacks,
		Business_Privacy: business,
		Expires_In_Days:  1,
	}

	attachment := entity.AttachmentTypeSendTemplate{
		Type:    "template",
		Payload: payload,
	}

	response := entity.MessangeSentTemplate{
		Recipient: entity.RecipientType{
			ID: senderId,
		},
		Message: entity.ResMessageTypeTemplate{
			Attachment: attachment,
		},
	}
	data, err := json.Marshal(response)
	if err != nil {
		log.Println("Marshal error: %s", err)
		return err
	}
	url_response := "https://graph.facebook.com/v7.0/me/messages"
	fb_token := os.Getenv("FACEBOOK_TOKEN")
	url_response = fmt.Sprintf("%s?access_token=%s", url_response, fb_token)
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
