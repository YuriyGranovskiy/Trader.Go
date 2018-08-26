package trade

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha512"
	"encoding/hex"
	"log"
	"net/http"
)

func GetAuthenticatedRequest(uri string, method string, requestBody []byte, key string, secret string) *http.Request {
	reader := bytes.NewReader(requestBody)

	req, err := http.NewRequest(method, uri, reader)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Key", key)
	req.Header.Set("Sign", computeHmac512(requestBody, secret))

	return req
}

func computeHmac512(message []byte, secret string) string {
	key := []byte(secret)
	h := hmac.New(sha512.New, key)
	h.Write([]byte(message))
	return hex.EncodeToString(h.Sum(nil))
}
