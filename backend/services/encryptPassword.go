package services

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string, salt string) (string, error) {
	var saltePassword = append([]byte(password), salt...)

	var hashPassword, err = bcrypt.GenerateFromPassword(saltePassword, bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashPassword), nil
}
