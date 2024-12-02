package auth

import (
	"database/sql"
	"fmt"

	"github.com/David-Alejandro-Jimenez/Pagina-futbol/internal/database"
)

func GetUser(username string) (bool, error) {
	var existingUser bool
	var err = database.DB.QueryRow("SELECT EXISTS(SELECT 1 FROM user_registration WHERE UserName = ?)", username).Scan(&existingUser)
	
		if err != nil {
			return false, fmt.Errorf("error consultando la base de datos %w", err)
		}
	return existingUser, nil
}

func GetHashPassword(username string) (string, error) {
	var hashPassword string
	var query = "SELECT Password FROM user_registration WHERE UserName= ?"
	var err = database.DB.QueryRow(query, username).Scan(&hashPassword)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", err
		}
		return "", err
	}

	return hashPassword, nil
}	

func SaveUser(userName, password string) error {
	var salt, errSalt = GenerateSalt()
	if errSalt != nil {
		return errSalt
	}
	var hash, err = HashPassword(password, salt)
	if err != nil {
		return err
	}

	_, err = database.DB.Exec("INSERT INTO user_registration (username, password, salt) VALUES (?, ?, ?)", userName, hash, salt)
	if err != nil {
		return err
	}
    return nil
}

func GetSalt(username string) (string, error) {
	var salt string
	var query = "SELECT salt FROM user_registration WHERE UserName= ?"
	var err  = database.DB.QueryRow(query, username).Scan(&salt)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", err
		}
		return "", err
	}

	return salt, nil
}


func GetUserID(username string) (int, error) {
	var id int
	var query = "SELECT id FROM user_registration WHERE UserName= ?"
	var err = database.DB.QueryRow(query, username).Scan(&id)
	if err != nil {
		if err == sql.ErrNoRows {
			return 0, err
		}
		return 0, err
	}
	return id, nil
}
