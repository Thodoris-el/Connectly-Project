package tests

import (
	"testing"
	"time"

	entity "github.com/Thodoris-el/Coonectly-Project/api/Entity"
)

func TestSaveTemplate(t *testing.T) {

	refreshTables()

	testTemplate := entity.Template{
		Title:       "How would you rate our product?",
		Placeholder: "Give additional feedback",
		Language:    "eng",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	savedTemplate, err := testTemplate.SaveTemplate(server.DB)

	if err != nil {
		t.Errorf("error saving the Template: %v\n", err)
		return
	}

	if !(testTemplate.Button_Title == savedTemplate.Button_Title && testTemplate.Language == savedTemplate.Language && testTemplate.Placeholder == savedTemplate.Placeholder && testTemplate.Subtitle == savedTemplate.Subtitle) {
		t.Errorf("wrong values")
	}
}

func TestFindAllTemplates(t *testing.T) {

	refreshTables()

	_, err := createTwoTemplates()
	testTemplate := entity.Template{}

	if err != nil {
		t.Errorf("error creating the two templates: %v\n", err)
		return
	}

	getT, err := testTemplate.FindAllTemplates(server.DB)
	if err != nil {
		t.Errorf("error fetching the Templates: %v\n", err)
		return
	}

	if len(*getT) != 2 {
		t.Errorf("wrong number of Templates")
	}
}

func TestFindByID(t *testing.T) {

	refreshTables()

	testTemplate, err := createATemplate()

	if err != nil {
		t.Errorf("error creating the Template: %v\n", err)
		return
	}

	getT, err := testTemplate.FindById(server.DB, 1)
	if err != nil {
		t.Errorf("error fetching the Review: %v\n", err)
		return
	}

	if !(testTemplate.Button_Title == getT.Button_Title && testTemplate.Language == getT.Language && testTemplate.Placeholder == getT.Placeholder && testTemplate.Subtitle == getT.Subtitle) {
		t.Errorf("wrong values")
	}
}

func TestFindbyLanguage(t *testing.T) {

	refreshTables()

	testTemplate, err := createATemplate()

	if err != nil {
		t.Errorf("error creating the Template: %v\n", err)
		return
	}

	getT, err := testTemplate.FindByLanguage(server.DB, "eng")
	if err != nil {
		t.Errorf("error fetching the Template: %v\n", err)
		return
	}

	if getT.Language != "eng" {
		t.Errorf("wrong Template")
	}
}

func TestUpdateTemplate(t *testing.T) {

	refreshTables()

	testTemplate, err := createATemplate()

	if err != nil {
		t.Errorf("error creating the Template: %v\n", err)
		return
	}

	testTemplate.Subtitle = "Buy"
	getT, err := testTemplate.UpdateTemplate(server.DB, 1)
	if err != nil {
		t.Errorf("error fetching the Template: %v\n", err)
		return
	}

	if !(testTemplate.Subtitle == getT.Subtitle) {
		t.Errorf("wrong values")
	}
}

func TestDeleteTemplate(t *testing.T) {

	refreshTables()

	testTemplate, err := createATemplate()

	if err != nil {
		t.Errorf("error creating the Template: %v\n", err)
		return
	}

	getT, err := testTemplate.DeleteTemplate(server.DB, 1)
	if err != nil {
		t.Errorf("error fetching the Template: %v\n", err)
		return
	}

	if getT != 1 {
		t.Errorf("wrong values")
	}
}
