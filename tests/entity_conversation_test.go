package tests

import (
	"testing"
	"time"

	entity "github.com/Thodoris-el/Coonectly-Project/api/Entity"
)

func TestSaveConversation(t *testing.T) {

	refreshTables()

	testConversation := entity.Conversation{
		Facebook_id: "6706612322695175",
		Type:        "Review",
		Stage:       "None",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	savedConversation, err := testConversation.SaveConversation(server.DB)

	if err != nil {
		t.Errorf("error saving the Conversation: %v\n", err)
		return
	}

	if !(testConversation.Facebook_id == savedConversation.Facebook_id && testConversation.Stage == savedConversation.Stage && testConversation.Type == savedConversation.Type) {
		t.Errorf("wrong values")
	}
}

func TestFindAllConversations(t *testing.T) {

	refreshTables()

	_, err := createTwoConversations()
	testConversation := entity.Conversation{}

	if err != nil {
		t.Errorf("error creating the Conversation: %v\n", err)
		return
	}

	getC, err := testConversation.FindAllConversations(server.DB)
	if err != nil {
		t.Errorf("error fetching the Conversation: %v\n", err)
		return
	}

	if len(*getC) != 2 {
		t.Errorf("wrong number of Conversations")
	}
}

func TestFindConversationByID(t *testing.T) {

	refreshTables()

	testConversation, err := createAConversation()

	if err != nil {
		t.Errorf("error creating the Conversation: %v\n", err)
		return
	}

	getC, err := testConversation.FindById(server.DB, 1)
	if err != nil {
		t.Errorf("error fetching the Conversation: %v\n", err)
		return
	}

	if !(testConversation.Facebook_id == getC.Facebook_id && testConversation.Stage == getC.Stage && testConversation.Type == getC.Type) {
		t.Errorf("wrong values")
	}
}

func TestFindConversationByCustomerID(t *testing.T) {

	refreshTables()

	testConversation, err := createAConversation()

	if err != nil {
		t.Errorf("error creating the Conversation: %v\n", err)
		return
	}

	getC, err := testConversation.FindByCustomerId(server.DB, "6706612322695175")
	if err != nil {
		t.Errorf("error fetching the Conversation: %v\n", err)
		return
	}

	if !(testConversation.Facebook_id == getC.Facebook_id && testConversation.Stage == getC.Stage && testConversation.Type == getC.Type) {
		t.Errorf("wrong values")
	}
}

func TestUpdateConversation(t *testing.T) {

	refreshTables()

	testConversation, err := createAConversation()

	if err != nil {
		t.Errorf("error creating the Conversation: %v\n", err)
		return
	}

	testConversation.Stage = "Buy"
	getC, err := testConversation.UpdateConversation(server.DB, 1)
	if err != nil {
		t.Errorf("error fetching the Conversation: %v\n", err)
		return
	}

	if !(testConversation.Facebook_id == getC.Facebook_id && testConversation.Stage == getC.Stage && testConversation.Type == getC.Type) {
		t.Errorf("wrong values")
	}
}

func TestDeleteConversation(t *testing.T) {

	refreshTables()

	testConversation, err := createAConversation()

	if err != nil {
		t.Errorf("error creating the Conversation: %v\n", err)
		return
	}

	getC, err := testConversation.DeleteConversation(server.DB, 1)
	if err != nil {
		t.Errorf("error fetching the Conversation: %v\n", err)
		return
	}

	if getC != 1 {
		t.Errorf("wrong values")
	}
}
