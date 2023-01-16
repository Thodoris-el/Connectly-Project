package tests

import (
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestVerifyWebhook(t *testing.T) {
	secret := os.Getenv("SECRET_TOKEN")
	var res []string
	res = append(res, "1158201444")
	res = append(res, "Bad token")
	requests := []http.Request{}
	request := httptest.NewRequest(http.MethodGet, "/webhookwebhook?hub.mode=subscribe&hub.challenge=1158201444&hub.verify_token="+secret, nil)
	requests = append(requests, *request)
	request = httptest.NewRequest(http.MethodGet, "/webhookwebhook?hub.mode=subscribe&hub.challenge=1158201444&hub.verify_token=good", nil)
	requests = append(requests, *request)
	for i, r := range requests {
		w := httptest.NewRecorder()
		server.VerifyWebhook(w, &r)
		result := w.Result()
		defer result.Body.Close()
		data, err := io.ReadAll(result.Body)
		if err != nil {
			t.Errorf("expected error to be nil got %v", err)
		}
		if string(data) != res[i] {
			t.Errorf("expected %v got %v", res[i], string(data))
		}
	}

}
