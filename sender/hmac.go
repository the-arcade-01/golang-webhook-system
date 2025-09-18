package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
)

// GenerateHMAC generates an HMAC hash using SHA256.
func GenerateHMAC(message []byte, secret string) string {
	h := hmac.New(sha256.New, []byte(secret))
	h.Write(message)
	return hex.EncodeToString(h.Sum(nil))
}

// GenerateSecretKey creates a secret key based on customer_id and webhook_url.
func GenerateSecretKey(customerID, webhookURL string) string {
	data := customerID + ":" + webhookURL
	hash := sha256.Sum256([]byte(data))
	return hex.EncodeToString(hash[:])
}
