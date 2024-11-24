package services

import (
	"fmt"
	"unicode"

	_ "github.com/go-sql-driver/mysql"
)

const (
	minUserNameLength = 5
	minPasswordLength = 10
)

func ValidateUserName(userName string) error {
	if len(userName) < minUserNameLength {
		return fmt.Errorf("no puedes ingresar un nombre que tenga menos de 5 caracteres")
	}
	return nil
}

func ValidatePassword(password string) error  {
	if len(password) < minPasswordLength {
		return fmt.Errorf("no puedes ingresar una contraseña que tenga menos de 10 caracteres")
	}

	var hasUppercase bool
	var hasDigit bool
	var hasSpecialCharacter bool
	for _, char := range password {
		if unicode.IsUpper(char) {
			hasUppercase = true
		}

		if unicode.IsDigit(char) {
			hasDigit = true
		}
		if unicode.IsPunct(char) || unicode.IsSymbol(char) {
			hasSpecialCharacter = true
		}

		if hasUppercase && hasDigit && hasSpecialCharacter{
			break
		}
	}

	if !hasUppercase {
		return fmt.Errorf("la contraseña debe tener al menos una letra mayuscula")
	}
	
	if !hasDigit {
		return fmt.Errorf("la contraseña debe tener al menos un numero")
	}

	if !hasSpecialCharacter {
		return fmt.Errorf("la contraseña debe tener algun caracter especial")
	}
	return nil
}

