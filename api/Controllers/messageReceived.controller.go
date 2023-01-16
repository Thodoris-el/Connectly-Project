package controllers

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"

	entity "github.com/Thodoris-el/Coonectly-Project/api/Entity"
	"gorm.io/gorm"
)

func (server *Server) HandleMessenger(resp http.ResponseWriter, request *http.Request) {
	//Read the body of the request received
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
		var fbFeed entity.MesFeedType

		for _, everyfbMess := range fbMess {
			sender = everyfbMess.Sender
			recipient = everyfbMess.Recipient
			message = everyfbMess.Message
			fbFeed = everyfbMess.Messaging_Feedback

			conversation := entity.Conversation{}
			new_conversation, err := conversation.FindByCustomerId(server.DB, sender.ID)

			if err != nil {
				if errors.Is(err, gorm.ErrRecordNotFound) {
					conversation.CreatedAt = time.Now()
					conversation.UpdatedAt = time.Now()
					conversation.Facebook_id = sender.ID
					new_conversation, err = conversation.SaveConversation(server.DB)
					if err != nil {
						log.Println("error while creating conversation")
						return
					}
				} else {
					log.Println("error while finding conversation")
					return
				}
			}

			if message.Attachments == nil && message.Text == "" && len(fbFeed.FeedbackScreens) > 0 {
				feedscreens := everyfbMess.Messaging_Feedback.FeedbackScreens
				for _, everyf := range feedscreens {

					if conversation.Stage == "Review" {
						score := everyf.Questions.Myquestion1.Payload
						err = server.AddReview(sender.ID, "", score)
						if err != nil {
							log.Println("error while creating review", err)
						}
						new_conversation.Stage = "None"
						_, err = new_conversation.UpdateConversation(server.DB, new_conversation.ID)
						if err != nil {
							log.Println("error while creating conversation")
							return
						}
						handleMessageWithoutQuickReply(sender.ID, "Thanks for the review!.")
					}
				}
			} else if message.Attachments == nil {

				if strings.Contains(message.Text, "Buy") {
					new_conversation.Stage = "Buy"
					_, err = new_conversation.UpdateConversation(server.DB, new_conversation.ID)
					if err != nil {
						log.Println("error while creating conversation")
						return
					}
					handleMessageWithQuickReply(sender.ID, "Are you sure?")

				} else if strings.Contains(message.Text, "Yes") {
					if new_conversation.Stage == "Buy" {
						new_conversation.Stage = "Review"
						_, err = new_conversation.UpdateConversation(server.DB, new_conversation.ID)
						if err != nil {
							log.Println("error while creating conversation")
							return
						}
						handleMessageWithoutQuickReply(sender.ID, "Purchased Done!")
						SendReviewTemplate(sender.ID)
					} else if new_conversation.Stage == "Review" {
						new_conversation.Stage = "None"
						_, err = new_conversation.UpdateConversation(server.DB, new_conversation.ID)
						if err != nil {
							log.Println("error while creating conversation")
							return
						}
						err = server.AddReview(sender.ID, message.Text, "-1")
						if err != nil {
							log.Println("error while creating database")
							return
						}
						handleMessageWithoutQuickReply(sender.ID, "Thanks for the review!.")

					} else {
						handleMessageWithoutQuickReply(sender.ID, "Didn't understand this! Tell us what product you want to buy.")
					}

				} else {
					if new_conversation.Stage == "Review" {
						new_conversation.Stage = "None"
						_, err = new_conversation.UpdateConversation(server.DB, new_conversation.ID)
						if err != nil {
							log.Println("error while creating conversation")
							return
						}
						err = server.AddReview(sender.ID, message.Text, "-1")
						if err != nil {
							log.Println("error while creating database")
							return
						}
						if err != nil {
							log.Println("error while creating database")
							return
						}
						handleMessageWithoutQuickReply(sender.ID, "Thanks for the review!.")
					} else if new_conversation.Stage == "Buy" {
						if strings.Contains(message.Text, "No") {
							new_conversation.Stage = "None"
							_, err = new_conversation.UpdateConversation(server.DB, new_conversation.ID)
							if err != nil {
								log.Println("error while creating conversation")
								return
							}
							handleMessageWithoutQuickReply(sender.ID, "Purchased Cancelled!")
						} else {
							new_conversation.Stage = "None"
							_, err = new_conversation.UpdateConversation(server.DB, new_conversation.ID)
							if err != nil {
								log.Println("error while creating conversation")
								return
							}
							handleMessageWithoutQuickReply(sender.ID, "Didn't understand this! Tell us what product you want to buy.")
						}
					} else {
						handleMessageWithoutQuickReply(sender.ID, "Didn't understand this! Tell us what product you want to buy")
					}
				}
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
