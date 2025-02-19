package app

import (
	"crypto/sha256"
	"encoding/base64"
)

func generateShortCode(url string) string {
	hash := sha256.Sum256([]byte(url))
	shortCode := base64.URLEncoding.EncodeToString(hash[:])
	return shortCode[:6]
}