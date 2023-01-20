package controllers

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"strings"
	"time"

	entity "github.com/Thodoris-el/Coonectly-Project/api/Entity"
	"gorm.io/gorm"
)

/*
This is the controller that handles our receiving messages
*/

func (server *Server) HandleMessengeQuickReply(message entity.MessageType, new_conversation entity.Conversation, sender string) (string, error) {
	//if customer replied that he bought the product
	if message.QuickReply.Payload == "Buy Product" {
		//change conversation stage to review
		new_conversation.Stage = "Review"
		_, err := new_conversation.UpdateConversation(server.DB, new_conversation.ID)
		if err != nil {
			log.Println("error while creating conversation")
			return "", err
		}

		//reply to customer with verifcation msg
		err = handleMessageWithoutQuickReply(sender, "Purchased Done!")
		if err != nil {
			log.Println(err.Error())
			return "", err
		}
		//find customer to see his language
		customer := entity.Customer{}
		customerGet, err := customer.FindByFacebookId(server.DB, sender)
		//if error getting the language just make use eng language for the template
		var lang string
		if err != nil {
			lang = "eng"
		} else {
			lang = customerGet.Language
		}
		//find the template of the review
		template := entity.Template{}
		templateGet, err := template.FindByLanguage(server.DB, lang)
		//if error getting the template just reply with message
		if err != nil {
			err = handleMessageWithoutQuickReply(sender, "Please write your review as a message!")
			if err != nil {
				log.Println(err.Error())
				return "", err
			}
			return "review", err
		}
		//send the review
		err = SendReviewTemplate(sender, templateGet)
		if err != nil {
			log.Println(err.Error())
			return "", err
		}
		return "review", err
		//if customer replied that he didnt buy the product
	} else if message.QuickReply.Payload == "Don't Buy Product" {
		//update conversation stage to none
		new_conversation.Stage = "None"
		_, err := new_conversation.UpdateConversation(server.DB, new_conversation.ID)
		if err != nil {
			log.Println("error while creating conversation")
			return "", err
		}
		//reply to customer that his purchase is cancelled
		err = handleMessageWithoutQuickReply(sender, "Purchase Cancelled!")
		if err != nil {
			log.Println(err.Error())
			return "", err
		}
		return "none", err
	}
	return "", errors.New("invalid quick reply answer")
}

func (server *Server) HandleMessageTemplate(everyfbMess entity.MessagingType, conversation, new_conversation entity.Conversation, sender string) (string, error) {
	feedscreens := everyfbMess.Messaging_Feedback.FeedbackScreens
	for _, everyf := range feedscreens {
		//Stage with prev message was the review template
		if conversation.Stage == "Review" {
			//Get values from the review template
			score := everyf.Questions.Myquestion1.Payload
			text := everyf.Questions.Myquestion1.FollowUp.Payload
			product := new_conversation.Product
			err := server.AddReview(sender, text, score, product)
			if err != nil {
				log.Println("error while creating review", err)
				return "", err
			}
			new_conversation.Stage = "None"
			new_conversation.Product = ""
			_, err = new_conversation.UpdateConversation(server.DB, new_conversation.ID)
			if err != nil {
				log.Println("error while creating conversation")
				return "", err
			}
			//send a thank you msg
			err = handleMessageWithoutQuickReply(sender, "Thanks for the review!.")
			if err != nil {
				log.Println(err.Error())
				return "", err
			}
			return "none", err
		}
	}
	return "", errors.New("invalid stage for review")
}

func (server *Server) SendTemplate(sender string) (string, error) {
	//Find customer that send the message
	customer := entity.Customer{}
	customerGet, err := customer.FindByFacebookId(server.DB, sender)
	//if error use eng as language for the template
	var lang string
	if err != nil {
		lang = "eng"
	} else {
		lang = customerGet.Language
	}
	//find the template
	template := entity.Template{}
	templateGet, err := template.FindByLanguage(server.DB, lang)
	//if error send it as a message
	if err != nil {
		err = handleMessageWithoutQuickReply(sender, "Please write your review as a message!")
		if err != nil {
			log.Println(err.Error())
			return "", err
		}
		return "review", err
	}
	err = SendReviewTemplate(sender, templateGet)
	if err != nil {
		log.Println(err.Error())
		return "", err
	}
	return "review", err
}

func (server *Server) HandleMessenger(resp http.ResponseWriter, request *http.Request) {
	//Read the body of the request received
	log.Println("Header: ", request.Header)
	body, err := io.ReadAll(request.Body)
	if err != nil {
		log.Printf("failed to read body: %v", err)
		http.Error(resp, err.Error(), http.StatusBadRequest)
		return
	}
	//Unmarshal request
	var facebookPost entity.FacebookMessage
	err = json.Unmarshal(body, &facebookPost)
	if err != nil {
		log.Println(err)
		http.Error(resp, err.Error(), http.StatusBadRequest)
		return
	}

	if facebookPost.Object != "page" {
		log.Println("No fb object page")
		http.Error(resp, errors.New("no fb object == page").Error(), http.StatusBadRequest)
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
			//if user message via quick answer
			if message.QuickReply.Payload != "" {
				log.Println(message)
				str, err := server.HandleMessengeQuickReply(message, *new_conversation, sender.ID)
				if err != nil {
					log.Println(err)
					return
				}
				resp.Header().Add("action", str)
				return
				//if user answers via review template
			} else if message.Attachments == nil && message.Text == "" && len(fbFeed.FeedbackScreens) > 0 {
				str, err := server.HandleMessageTemplate(everyfbMess, conversation, *new_conversation, sender.ID)
				if err != nil {
					log.Println(err)
				}
				resp.Header().Add("action", str)
				return
				//if user answers via text
			} else if message.Attachments == nil {

				//Buy is a trigger word
				if strings.HasPrefix(message.Text, "Buy a ") {
					new_conversation.Stage = "Buy"
					prod := strings.Split(message.Text, "Buy a ")
					new_conversation.Product = prod[len(prod)-1]
					_, err = new_conversation.UpdateConversation(server.DB, new_conversation.ID)
					if err != nil {
						log.Println("error while creating conversation")
						return
					}
					err = handleMessageWithQuickReply(sender.ID, "Are you sure?")
					if err != nil {
						log.Println(err.Error())
						return
					}
					resp.Header().Add("action", "buy")
					return

				} else if strings.Contains(message.Text, "Yes") {

					if new_conversation.Stage == "Buy" {
						new_conversation.Stage = "Review"
						_, err = new_conversation.UpdateConversation(server.DB, new_conversation.ID)
						if err != nil {
							log.Println("error while creating conversation")
							return
						}
						err = handleMessageWithoutQuickReply(sender.ID, "Purchased Done!")
						if err != nil {
							log.Println(err.Error())
							return
						}
						str, err := server.SendTemplate(sender.ID)
						if err != nil {
							log.Println(err)
							return
						}
						resp.Header().Add("action", str)

					} else if new_conversation.Stage == "Review" {
						err = server.AddReview(sender.ID, message.Text, "-1", new_conversation.Product)
						if err != nil {
							log.Println("error while creating database")
							return
						}
						new_conversation.Stage = "None"
						new_conversation.Product = ""
						_, err = new_conversation.UpdateConversation(server.DB, new_conversation.ID)
						if err != nil {
							log.Println("error while creating conversation")
							return
						}
						err = handleMessageWithoutQuickReply(sender.ID, "Thanks for the review!.")
						if err != nil {
							log.Println(err.Error())
							return
						}
						resp.Header().Add("action", "none")

					} else {
						err = handleMessageWithoutQuickReply(sender.ID, "Didn't understand this! Tell us what product you want to buy.")
						if err != nil {
							log.Println(err.Error())
							return
						}
						resp.Header().Add("action", "none")
						return
					}

				} else {

					if new_conversation.Stage == "Review" {
						err = server.AddReview(sender.ID, message.Text, "-1", new_conversation.Product)
						if err != nil {
							log.Println("error while creating database")
							return
						}
						new_conversation.Stage = "None"
						new_conversation.Product = ""
						_, err = new_conversation.UpdateConversation(server.DB, new_conversation.ID)
						if err != nil {
							log.Println("error while creating conversation")
							return
						}
						if err != nil {
							log.Println("error while creating database")
							return
						}
						err = handleMessageWithoutQuickReply(sender.ID, "Thanks for the review!.")
						if err != nil {
							log.Println(err.Error())
							return
						}
						resp.Header().Add("action", "none")
						return

					} else if new_conversation.Stage == "Buy" {

						if strings.Contains(message.Text, "No") {
							new_conversation.Stage = "None"
							new_conversation.Product = ""
							_, err = new_conversation.UpdateConversation(server.DB, new_conversation.ID)
							if err != nil {
								log.Println("error while creating conversation")
								return
							}
							err = handleMessageWithoutQuickReply(sender.ID, "Purchased Cancelled!")
							if err != nil {
								log.Println(err.Error())
								return
							}
							resp.Header().Add("action", "none")
							return

						} else {
							new_conversation.Stage = "None"
							new_conversation.Product = ""
							_, err = new_conversation.UpdateConversation(server.DB, new_conversation.ID)
							if err != nil {
								log.Println("error while creating conversation")
								return
							}
							err = handleMessageWithoutQuickReply(sender.ID, "Didn't understand this! Tell us what product you want to buy.")
							if err != nil {
								log.Println(err.Error())
								return
							}
							resp.Header().Add("action", "none")
							return
						}

					} else {
						err = handleMessageWithoutQuickReply(sender.ID, "Didn't understand this! Tell us what product you want to buy")
						if err != nil {
							log.Println(err.Error())
							return
						}
						resp.Header().Add("action", "none")
						return
					}
				}
			} else {
				attachment := message.Attachments
				for _, everyAttachment := range attachment {
					switch everyAttachment.Type {
					case "template":
						new_conversation.Stage = "None"
						new_conversation.Product = ""
						_, err = new_conversation.UpdateConversation(server.DB, new_conversation.ID)
						if err != nil {
							log.Println("error while creating conversation")
							return
						}
						log.Println(sender, recipient, everyAttachment.Payload.Product.Title)
						err = handleMessageWithoutQuickReply(sender.ID, "Didn't understand this! Tell us what product you want to buy")
						if err != nil {
							log.Println(err.Error())
							return
						}
						resp.Header().Add("action", "none")
						return
					default:
						new_conversation.Stage = "None"
						new_conversation.Product = ""
						_, err = new_conversation.UpdateConversation(server.DB, new_conversation.ID)
						if err != nil {
							log.Println("error while creating conversation")
							return
						}
						log.Println(sender, recipient, "image: ", everyAttachment.Payload.URL)
						err = handleMessageWithoutQuickReply(sender.ID, "Didn't understand this! Tell us what product you want to buy")
						if err != nil {
							log.Println(err.Error())
							return
						}
						resp.Header().Add("action", "none")
						return
					}
				}
			}
		}
	}
}
