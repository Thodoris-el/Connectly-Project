package tests

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandleMessanger(t *testing.T) {

	refreshTables()
	_, err := createACustomer()
	if err != nil {
		log.Fatal(err)
	}

	resmsg := []string{"buy", "review", "buy", "none", "none"}
	msg := []string{"Buy a car", "Yes", "Buy a Bike", "No", "hello"}
	for i, fbmess := range createMessages(msg) {
		body, _ := json.Marshal(fbmess)
		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodPost, "/webhook", bytes.NewReader(body))
		r.Header.Add("Content-Type", "application/json")
		r.Header.Add("X-Hub-Signature", "sha1=a502207038bb5692cb665062846716fbc7d7951a")
		r.Header.Add("X-Hub-Signature", "sha256=691990bf36efdb92312b367e252203d53e6fd4ba19666e74081bdadd9770188d")
		server.HandleMessenger(w, r)
		res := w.Result()

		if res.Header.Get("action") != resmsg[i] {
			t.Errorf("expected %v got %v", resmsg[i], res.Header.Get("action"))
		}
	}

	resmsg = []string{"buy", "review", "none"}
	msg = []string{"Buy a car", "Yes"}
	fbmessages := createMessages(msg)
	fbmessages = append(fbmessages, createMessageReview())
	for i, fbmess := range fbmessages {
		body, _ := json.Marshal(fbmess)
		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodPost, "/webhook", bytes.NewReader(body))
		r.Header.Add("Content-Type", "application/json")
		r.Header.Add("X-Hub-Signature", "sha1=a502207038bb5692cb665062846716fbc7d7951a")
		r.Header.Add("X-Hub-Signature", "sha256=691990bf36efdb92312b367e252203d53e6fd4ba19666e74081bdadd9770188d")
		server.HandleMessenger(w, r)
		res := w.Result()

		if res.Header.Get("action") != resmsg[i] {
			t.Errorf("expected %v got %v", resmsg[i], res.Header.Get("action"))
		}
	}
}
