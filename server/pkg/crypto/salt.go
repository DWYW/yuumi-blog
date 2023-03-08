package crypto

import (
	"crypto/rand"
	"encoding/base64"
)

func GetSlat(len int) (string, error) {
	key := make([]byte, len)
	_, err := rand.Read(key[:])

	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(key), nil
}
