package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	entity "github.com/Thodoris-el/Coonectly-Project/api/Entity"
	"github.com/gorilla/mux"
)

func TestCreateTemplate(t *testing.T) {

	refreshTables()

	samples := []string{
		`{"title": "new1","placeholder": "p1","language": "eng"}`,
		`{"title": "new2","placeholder": "p2","language": "german"}`,
		`{"title": "new3","placeholder": "p3","language": "french"}`,
		`{"title": "new4","placeholder": "p4","language": "eng"}`,
	}
	answers := []string{
		"new1",
		"new2",
		"new3",
		"new4",
	}

	for i, tmp := range samples {
		req, err := http.NewRequest("POST", "/template", bytes.NewBufferString(tmp))
		if err != nil {
			t.Errorf("error: %v\n", err)
		}
		recorded := httptest.NewRecorder()
		handler := http.HandlerFunc(server.CreateTemplate)
		handler.ServeHTTP(recorded, req)

		resp := make(map[string]interface{})
		err = json.Unmarshal(recorded.Body.Bytes(), &resp)
		if err != nil {
			t.Errorf("cant convert json: %v\n", err)
		}
		if resp["title"] != answers[i] {
			t.Errorf("wrong values")
		}
	}
}

func TestGetTemplate(t *testing.T) {

	refreshTables()

	_, err := createTwoTemplates()
	if err != nil {
		t.Errorf("error creating  templates")
	}

	req, err := http.NewRequest("GET", "/template", nil)
	if err != nil {
		t.Errorf("error: %v\n", err)
	}
	recorded := httptest.NewRecorder()
	handler := http.HandlerFunc(server.GetTemplate)
	handler.ServeHTTP(recorded, req)

	var templates []entity.Template
	err = json.Unmarshal(recorded.Body.Bytes(), &templates)
	if err != nil {
		t.Errorf("error while unmarshal")
	}
	if len(templates) != 2 {
		t.Errorf("wrong number of templates")
	}
}

func TestGetTemplateById(t *testing.T) {

	refreshTables()

	_, err := createATemplate()
	if err != nil {
		t.Errorf("error creating a templates")
	}

	req, err := http.NewRequest("GET", "/template", nil)
	if err != nil {
		t.Errorf("error: %v\n", err)
	}
	req = mux.SetURLVars(req, map[string]string{"id": "1"})
	recorded := httptest.NewRecorder()
	handler := http.HandlerFunc(server.GetTemplateById)
	handler.ServeHTTP(recorded, req)

	responseMap := make(map[string]interface{})
	err = json.Unmarshal(recorded.Body.Bytes(), &responseMap)
	if err != nil {
		t.Errorf("error while unmarshal")
	}
	if responseMap["title"] != "How would you rate our product?" {
		t.Errorf("wrong value")
	}
}

func TestGetTemplateByLanguage(t *testing.T) {

	refreshTables()

	_, err := createATemplate()
	if err != nil {
		t.Errorf("error creating a template")
	}

	req, err := http.NewRequest("GET", "/template/language", nil)
	if err != nil {
		t.Errorf("error: %v\n", err)
	}
	req = mux.SetURLVars(req, map[string]string{"language": "eng"})
	recorded := httptest.NewRecorder()
	handler := http.HandlerFunc(server.GetTemplateByLanguage)
	handler.ServeHTTP(recorded, req)

	responseMap := make(map[string]interface{})
	err = json.Unmarshal(recorded.Body.Bytes(), &responseMap)
	if err != nil {
		t.Errorf("error while unmarshal")
	}
	if responseMap["title"] != "How would you rate our product?" {
		t.Errorf("wrong value")
	}
}

func TestUpdateTemplateC(t *testing.T) {

	refreshTables()

	_, err := createATemplate()
	if err != nil {
		t.Errorf("error creating a template")
	}
	sample := `{"title": "Buy"}`
	req, err := http.NewRequest("PUT", "/template", bytes.NewBufferString(sample))
	if err != nil {
		t.Errorf("error: %v\n", err)
	}

	req = mux.SetURLVars(req, map[string]string{"id": "1"})
	recorded := httptest.NewRecorder()
	handler := http.HandlerFunc(server.UpdateTemplate)
	handler.ServeHTTP(recorded, req)

	responseMap := make(map[string]interface{})
	err = json.Unmarshal(recorded.Body.Bytes(), &responseMap)
	if err != nil {
		t.Errorf("error while unmarshal")
	}
	if responseMap["title"] != "Buy" {
		t.Errorf("wrong value")
	}
}

func TestDeleteTemplatec(t *testing.T) {

	refreshTables()

	_, err := createATemplate()
	if err != nil {
		t.Errorf("error creating a template")
	}

	req, err := http.NewRequest("DELETE", "/template", nil)
	if err != nil {
		t.Errorf("error: %v\n", err)
	}
	req = mux.SetURLVars(req, map[string]string{"id": "1"})
	recorded := httptest.NewRecorder()
	handler := http.HandlerFunc(server.DeleteTemplate)
	handler.ServeHTTP(recorded, req)

	var templates entity.Template
	getT, err := templates.FindAllTemplates(server.DB)
	if err != nil {
		t.Errorf("error while getting templates")
	}
	if len(*getT) != 0 {
		t.Errorf("wrong number of templates")
	}
}
