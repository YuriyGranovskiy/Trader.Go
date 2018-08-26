package trade

import (
	"crypto/hmac"
	"crypto/sha512"
	"encoding/hex"
)

const AuthKey = ""
const AuthSecret = ""

func computeHmac512(message string, secret string) string {
	key := []byte(secret)
	h := hmac.New(sha512.New, key)
	h.Write([]byte(message))
	return hex.EncodeToString(h.Sum(nil))
}
