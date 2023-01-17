package utilities

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
)

func createHash(secret, payload []byte) string {

	tmp := hmac.New(sha256.New, secret)
	tmp.Write(payload)
	base64Hash := base64.StdEncoding.EncodeToString(tmp.Sum(nil))
	return base64Hash
}

func verifyHash(secret_token, verifyFB string, payload []byte) bool {

	tmp := createHash([]byte(secret_token), payload)
	return hmac.Equal([]byte(verifyFB), []byte(tmp))
}
