package database

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/David-Alejandro-Jimenez/Pagina-futbol/services"
)

func RecoverStoredSalt(username string) (string, error) {
	var salt string
	var query = "SELECT salt FROM user_registration WHERE UserName= ?"
	var err  = DataBase.QueryRow(query, username).Scan(&salt)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", err
		}
		return "", err
	}

	return salt, nil
}

func RecoverStoredHashPassword(username string) (string, error) {
	var hashPassword string
	var query = "SELECT Password FROM user_registration WHERE UserName= ?"
	var err = DataBase.QueryRow(query, username).Scan(&hashPassword)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", err
		}
		return "", err
	}

	return hashPassword, nil
}

func ValidateExistingUsers(username string) (bool, error) {
	var existingUser bool
	var err = DataBase.QueryRow("SELECT EXISTS(SELECT 1 FROM user_registration WHERE UserName = ?)", username).Scan(&existingUser)
	
		if err != nil {
			log.Printf("Error al consultar si el usuario existe: %v", err)
			return false, fmt.Errorf("error consultando la base de datos %w", err)
		}
	return existingUser, nil
}

func SaveUser(userName, password string) error {
	var salt, errSalt = services.GenerateSalt()
	if errSalt != nil {
		return errSalt
	}
	var hash, err = services.HashPassword(password, salt)
	if err != nil {
		return err
	}

	_, err = DataBase.Exec("INSERT INTO user_registration (username, password, salt) VALUES (?, ?, ?)", userName, hash, salt)
	if err != nil {
		return err
	}
    return nil
}