package database

import (
	"fmt"
	"log"

	"github.com/David-Alejandro-Jimenez/Pagina-futbol/services"
)

func ValidateExistingUsers(username string) (bool, error) {
	var existingUser bool
	var err = DataBase.QueryRow("SELECT EXISTS(SELECT 1 FROM users WHERE UserName=?)", username).Scan(&existingUser)
	
		if err != nil {
			log.Printf("Error al consultar si el usuario existe: %v", err)
			return false, fmt.Errorf("error consultando la base de datos %w", err)
		}
	return existingUser, nil
}

func SaveUser(userName, password string) error {
	var hash, err = services.HashPassword(password)
	if err != nil {
		return err
	}

	_, err = DataBase.Exec("INSERT INTO users (username, password) VALUES (?, ?)", userName, hash)
    return err
}