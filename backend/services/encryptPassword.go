package services

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {
	var hashPassword, err = bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hashPassword), err
}
