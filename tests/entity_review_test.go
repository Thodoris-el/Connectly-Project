package tests

import (
	"testing"
	"time"

	entity "github.com/Thodoris-el/Coonectly-Project/api/Entity"
)

func TestSaveReview(t *testing.T) {

	refreshTables()

	testReview := entity.Review{
		Customer_id: "6706612322695175",
		Text:        "None",
		Score:       1,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	savedReview, err := testReview.SaveReview(server.DB)

	if err != nil {
		t.Errorf("error saving the Review: %v\n", err)
		return
	}

	if !(testReview.Customer_id == savedReview.Customer_id && testReview.Text == savedReview.Text && testReview.Score == savedReview.Score) {
		t.Errorf("wrong values")
	}
}

func TestFindAllReviews(t *testing.T) {

	refreshTables()

	_, err := createTwoReviews()
	testReview := entity.Review{}

	if err != nil {
		t.Errorf("error creating the Review: %v\n", err)
		return
	}

	getR, err := testReview.FindAllReviews(server.DB)
	if err != nil {
		t.Errorf("error fetching the Review: %v\n", err)
		return
	}

	if len(*getR) != 2 {
		t.Errorf("wrong number of Reviews")
	}
}

func TestFindReviewByID(t *testing.T) {

	refreshTables()

	testReview, err := createAReview()

	if err != nil {
		t.Errorf("error creating the Review: %v\n", err)
		return
	}

	getR, err := testReview.FindById(server.DB, 1)
	if err != nil {
		t.Errorf("error fetching the Review: %v\n", err)
		return
	}

	if !(testReview.Customer_id == getR.Customer_id && testReview.Text == getR.Text && testReview.Score == getR.Score) {
		t.Errorf("wrong values")
	}
}

func TestFindReviewByCustomerID(t *testing.T) {

	refreshTables()

	testReview, err := createAReview()

	if err != nil {
		t.Errorf("error creating the Review: %v\n", err)
		return
	}

	getR, err := testReview.FindByCustomerId(server.DB, "6706612322695175")
	if err != nil {
		t.Errorf("error fetching the Review: %v\n", err)
		return
	}

	if len(*getR) != 1 {
		t.Errorf("wrong number of Reviews")
	}
}

func TestUpdateReview(t *testing.T) {

	refreshTables()

	testReview, err := createAReview()

	if err != nil {
		t.Errorf("error creating the Review: %v\n", err)
		return
	}

	testReview.Text = "Buy"
	getR, err := testReview.UpdateReview(server.DB, 1)
	if err != nil {
		t.Errorf("error fetching the Review: %v\n", err)
		return
	}

	if !(testReview.Customer_id == getR.Customer_id && testReview.Text == getR.Text && testReview.Score == getR.Score) {
		t.Errorf("wrong values")
	}
}

func TestDeleteReview(t *testing.T) {

	refreshTables()

	testReview, err := createAReview()

	if err != nil {
		t.Errorf("error creating the Review: %v\n", err)
		return
	}

	getR, err := testReview.DeleteReview(server.DB, 1)
	if err != nil {
		t.Errorf("error fetching the Review: %v\n", err)
		return
	}

	if getR != 1 {
		t.Errorf("wrong values")
	}
}
