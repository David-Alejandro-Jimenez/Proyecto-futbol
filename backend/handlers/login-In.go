package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/David-Alejandro-Jimenez/Pagina-futbol/database"
	"github.com/David-Alejandro-Jimenez/Pagina-futbol/models"
	"github.com/David-Alejandro-Jimenez/Pagina-futbol/services"
	"golang.org/x/crypto/bcrypt"
)

func LoginIn(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}

	var application models.Account
	var errapplication = json.NewDecoder(r.Body).Decode(&application)
	if errapplication != nil {
		http.Error(w, "Solicitud Invalida", http.StatusBadRequest)
		return
	}

	var errValidateUserName = services.ValidateUserName(application.UserName)
	if errValidateUserName != nil {
		http.Error(w, errValidateUserName.Error(), http.StatusBadRequest)
		return
	}

	var exists, errValidateExistingUser = database.ValidateExistingUsers(application.UserName)
	if !exists {
		http.Error(w, "Usuario o contraseña incorrectos", http.StatusUnauthorized)
		return
	}
	if errValidateExistingUser != nil {
		log.Printf("Error verificando usuario existente: %v", errValidateExistingUser)
		http.Error(w, "Error en el servidor al validar el usuario", http.StatusInternalServerError)
		return
	} 
	var salt, errSalt = database.RecoverStoredSalt(application.UserName)
	if errSalt != nil {
		http.Error(w, "Error en el servidor al recuperar el salt", http.StatusInternalServerError)
		return
	}
	
	var storeHash, errStoreHash = database.RecoverStoredHashPassword(application.UserName)
	if errStoreHash != nil {
		http.Error(w, "Error en el servidor al recuperar el hash", http.StatusInternalServerError)
		return
	}

	var passwordWithSalt = append([]byte(application.Password), salt...)
	var errComparePassword = bcrypt.CompareHashAndPassword([]byte(storeHash), passwordWithSalt)
	if errComparePassword != nil {
		http.Error(w, "Usuario o contraseña incorrectos", http.StatusUnauthorized)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Inicio de sesión exitoso"))
}