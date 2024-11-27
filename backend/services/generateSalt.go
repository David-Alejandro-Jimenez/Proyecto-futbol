package services

import (
	"crypto/rand"
	"encoding/base64"
)

func GenerateSalt() (string, error) {
	var salt = make([]byte, 16)
	var _, err = rand.Read(salt)
	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(salt), nil
}