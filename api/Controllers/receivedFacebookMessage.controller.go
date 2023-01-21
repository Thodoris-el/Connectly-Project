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
Type of Conversations that is supported right now:
	-> Review:
		->A customer buys a product and we send him a review template
		->Stages ot this conversation type:
			->Stage None -> we are waiting for the trigger phrase: Buy a <product>
			->Stage Buy -> we have heard the trigger phrase and we have sent a verification msg
			->Stage Review -> we get the verification from the customer and we send him the review template
*/
//handler that starts a review conversation type
func (server *Server) StartReviewConversation(sender, message string, new_conversation entity.Conversation) error {
	var err error
	//Start a conversation that is a review type
	new_conversation.Type = "Review"
	//stage of the review -> buy -> we send a verification msg to the customer
	new_conversation.Stage = "Buy"
	prod := strings.Split(message, "Buy a ")
	new_conversation.Product = prod[len(prod)-1] //save the product that the customer wants to buy
	_, err = new_conversation.UpdateConversation(server.DB, new_conversation.ID)
	if err != nil {
		log.Println("error while creating conversation")
		return err
	}
	//send verification msg with quick reply
	err = handleMessageWithQuickReplyReview(sender, "Are you sure?", new_conversation.Type)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	return nil
}

//Handler for the review type conversation
func (server *Server) HandleReviewTypeConersation(sender, message string, new_conversation entity.Conversation) (string, error) {
	var err error
	switch {
	//we have sent the verification msg and the customer answered yes
	case new_conversation.Stage == "Buy" && strings.Contains(strings.ToLower(message), "yes"):
		new_conversation.Stage = "Review" //change the conversation stage
		_, err = new_conversation.UpdateConversation(server.DB, new_conversation.ID)
		if err != nil {
			log.Println("error while creating conversation")
			return "", err
		}
		//send a verification msg
		err = handleMessageWithoutQuickReply(sender, "Your purchase has been confirmed!")
		if err != nil {
			log.Println(err.Error())
			return "", err
		}
		//send the review template
		str, err := server.SendTemplate(sender, new_conversation.Product)
		if err != nil {
			log.Println(err)
			return "", err
		}
		return str, err
	//We sent the verification msg and the customer answered yes
	case new_conversation.Stage == "Buy" && strings.Contains(strings.ToLower(message), "no"):
		//change conversation type, stage and product to none -> reset conversation
		new_conversation.Type = "None"
		new_conversation.Product = ""
		new_conversation.Stage = "None"
		_, err = new_conversation.UpdateConversation(server.DB, new_conversation.ID)
		if err != nil {
			log.Println("error while creating conversation")
			return "", err
		}
		//send a msg to the customer that informs him about the cancellation
		err = handleMessageWithoutQuickReply(sender, "Your purchase has been canceled!")
		if err != nil {
			log.Println(err.Error())
			return "", err
		}
		return "none", err
	//Customer sent us an invalid answer
	default:
		//reset the conversation
		new_conversation.Type = "None"
		new_conversation.Product = ""
		new_conversation.Stage = "None"
		_, err = new_conversation.UpdateConversation(server.DB, new_conversation.ID)
		if err != nil {
			log.Println("error while creating conversation")
			return "", err
		}
		//sent error msg
		err = handleMessageWithoutQuickReply(sender, "Something went wrong, please try again!")
		if err != nil {
			log.Println(err.Error())
			return "", err
		}
		return "none", err
	}
}

//this function handles the quick replies
func (server *Server) HandleReceivedFacebookMessageQuickReply(message entity.MessageType, new_conversation entity.Conversation, sender string) (string, error) {
	switch {
	//customer in review type
	case new_conversation.Type == "Review":
		var err error
		var str string
		switch message.QuickReply.Payload {
		//answered yes-> to buy a product
		case "Buy Product":
			str, err = server.HandleReviewTypeConersation(sender, "yes", new_conversation)
			if err != nil {
				log.Println(err)
				return "", err
			}
			return str, err
		//answered no -> do not buy a product
		case "Don't Buy Product":
			str, err = server.HandleReviewTypeConersation(sender, "no", new_conversation)
			if err != nil {
				log.Println(err)
				return "", err
			}
			return str, err
		//all other cases -> error
		default:
			//reset the conversation if needed
			if new_conversation.Type != "None" {
				new_conversation.Type = "None"
				new_conversation.Stage = "None"
				new_conversation.Product = ""
				_, err := new_conversation.UpdateConversation(server.DB, new_conversation.ID)
				if err != nil {
					log.Println("error while creating conversation")
					return "", err
				}
			}
			//reply to customer that something went wrong
			err := handleMessageWithoutQuickReply(sender, "Something went wrong, please try again!")
			if err != nil {
				log.Println(err.Error())
				return "", err
			}
			return "none", err
		}
	}
	return "none", nil
}

//this function handles the review template that we received from the customer
func (server *Server) HandleReceivedFacebookMessageTemplate(everyfbMess entity.MessagingType, conversation, new_conversation entity.Conversation, sender string) (string, error) {
	feedscreens := everyfbMess.Messaging_Feedback.FeedbackScreens
	for _, everyf := range feedscreens {
		switch {
		//ifright conversation type and stage -> take review
		case (conversation.Type == "Review" && conversation.Stage == "Review"):
			//Get values from the review template
			score := everyf.Questions.Myquestion1.Payload
			text := everyf.Questions.Myquestion1.FollowUp.Payload
			product := new_conversation.Product
			err := server.AddReview(sender, text, score, product)
			if err != nil {
				log.Println("error while creating review", err)
				return "", err
			}
			new_conversation.Type = "None"
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
		default:
			//reset the conversation if needed
			if conversation.Type != "None" {
				new_conversation.Type = "None"
				new_conversation.Stage = "None"
				new_conversation.Type = ""
				_, err := new_conversation.UpdateConversation(server.DB, new_conversation.ID)
				if err != nil {
					log.Println("error while creating conversation")
					return "", err
				}
			}
			//reply to customer that something went wrong
			err := handleMessageWithoutQuickReply(sender, "Something went wrong, please try again!")
			if err != nil {
				log.Println(err.Error())
				return "", err
			}

			return "none", err
		}
	}
	return "", errors.New("no template found")
}

func (server *Server) HandleReceivedFacebookMessage(resp http.ResponseWriter, request *http.Request) {
	//Read the body of the received facebook message
	body, err := io.ReadAll(request.Body)
	//if error while reading body
	if err != nil {
		log.Printf("failed to read body: %v", err)
		http.Error(resp, err.Error(), http.StatusBadRequest)
		return
	}

	//Unmarshal the facebook message
	var facebookPost entity.FacebookMessage
	err = json.Unmarshal(body, &facebookPost)
	//if error while unmarshal
	if err != nil {
		log.Println(err)
		http.Error(resp, err.Error(), http.StatusBadRequest)
		return
	}

	//muct contain page as an object -> otherwise error
	if facebookPost.Object != "page" {
		log.Println("No fb object page")
		http.Error(resp, errors.New("no fb object == page").Error(), http.StatusBadRequest)
	}

	fbEntry := facebookPost.Entry
	var fbMess []entity.MessagingType

	for _, everyEntry := range fbEntry {
		fbMess = everyEntry.Messaging
		var sender entity.SenderType
		var message entity.MessageType
		var fbFeed entity.MesFeedType

		for _, everyfbMess := range fbMess {
			sender = everyfbMess.Sender             //the sender of the msg -> contains the facebook id of the sender
			message = everyfbMess.Message           //the msg -> contains the text or an attachment or a quick reply
			fbFeed = everyfbMess.Messaging_Feedback // contains the review template feedback

			//find the conversation with the customer to see what conversation the customer has started and what we are supposed to answer
			conversation := entity.Conversation{}
			new_conversation, err := conversation.FindByCustomerId(server.DB, sender.ID)
			//if error finding the coversation
			if err != nil {
				//Check if the conversation did not exist -> if true -> create one
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

			switch {
			//If customer answered via Quick Answer
			case message.QuickReply.Payload != "":
				stage_of_conversation, err := server.HandleReceivedFacebookMessageQuickReply(message, *new_conversation, sender.ID)
				if err != nil {
					log.Println(err)
					return
				}
				resp.Header().Add("action", stage_of_conversation)
				return
			//If customer answered via Review Template -> the customer filled the review template
			case len(fbFeed.FeedbackScreens) > 0:
				str, err := server.HandleReceivedFacebookMessageTemplate(everyfbMess, conversation, *new_conversation, sender.ID)
				if err != nil {
					log.Println(err)
				}
				resp.Header().Add("action", str)
				return
			//if their are no attachments and customer answered via text
			case message.Attachments == nil:
				switch {
				//Trigger Word Buy -> customer sent a message like this: Buy a <product>
				case strings.HasPrefix(message.Text, "Buy a "):
					//Start a conversation that is a review type
					err = server.StartReviewConversation(sender.ID, message.Text, *new_conversation)
					if err != nil {
						log.Println(err.Error())
						return
					}
					resp.Header().Add("action", "buy")
					return
				//the customer has started a review conversation -> handle the review conversation type
				case new_conversation.Type == "Review":
					str, err := server.HandleReviewTypeConersation(sender.ID, message.Text, *new_conversation)
					if err != nil {
						log.Println(err.Error())
						return
					}
					resp.Header().Add("action", str)
				//if none of the above send error msg
				default:
					//reset the conversation if needed
					if new_conversation.Type != "None" {
						new_conversation.Type = "None"
						new_conversation.Product = ""
						new_conversation.Stage = "None"
						_, err = new_conversation.UpdateConversation(server.DB, new_conversation.ID)
						if err != nil {
							log.Println("error while creating conversation")
							return
						}
					}
					err = handleMessageWithoutQuickReply(sender.ID, "Something went wrong, please try again!")
					if err != nil {
						log.Println(err.Error())
						return
					}
					resp.Header().Add("action", "none")
				}
			//if nothing of the above -> ex. customer send us an attachment
			default:
				//reset the conversation only if it is needed
				if conversation.Type != "None" {
					new_conversation.Type = "None"
					new_conversation.Product = ""
					new_conversation.Stage = "None"
					_, err = new_conversation.UpdateConversation(server.DB, new_conversation.ID)
					if err != nil {
						log.Println("error while creating conversation")
						return
					}
				}
				err = handleMessageWithoutQuickReply(sender.ID, "An error occurred, please try again!")
				if err != nil {
					log.Println(err.Error())
					return
				}
				resp.Header().Add("action", "none")
			}
		}
	}
}
