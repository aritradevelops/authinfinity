package crypto

import (
	"crypto/rand"
	"encoding/hex"
)

// GenerateHash generates a random bytes of the len given.
// and returns it in hex encoded string
func GenerateHash(length int) (string, error) {
	secretBytes := make([]byte, length)
	_, err := rand.Read(secretBytes)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(secretBytes), nil
}
