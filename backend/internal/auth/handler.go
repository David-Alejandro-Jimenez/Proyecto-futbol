package auth

import (
	"encoding/json"
	"net/http"

	"github.com/David-Alejandro-Jimenez/Pagina-futbol/models"
	"github.com/David-Alejandro-Jimenez/Pagina-futbol/services"
	"golang.org/x/crypto/bcrypt"
)

var err error

func RegisterNewAccount(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}

	var application models.Account
	err = json.NewDecoder(r.Body).Decode(&application)
	if err != nil {
		http.Error(w, "Solicitud Invalida", http.StatusBadRequest)
		return
	}

	err = ValidateUserName(application.UserName)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	
	err = ValidatePassword(application.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var exists, errGetUser = GetUser(application.UserName) 
	if errGetUser != nil {
    	http.Error(w, "Error en el servidor al validar el usuario", http.StatusInternalServerError)
   		return
	}

	if exists {
		http.Error(w, "El nombre de usuario ya existe", http.StatusConflict)
		return
	}

	err = SaveUser(application.UserName, application.Password)
	if err != nil {
		http.Error(w, "Error al guardar el usuario en la base de datos", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Usuario creado exitosamente"})
}


func LoginIn(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}

	var application models.Account
	err = json.NewDecoder(r.Body).Decode(&application)
	if err != nil {
		http.Error(w, "Solicitud Invalida", http.StatusBadRequest)
		return
	}

	err = ValidateUserName(application.UserName)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var exists, errGetUser = GetUser(application.UserName)
	if !exists {
		http.Error(w, "Usuario o contraseña incorrectos", http.StatusUnauthorized)
		return
	}
	if errGetUser != nil {
		http.Error(w, "Error en el servidor al validar el usuario", http.StatusInternalServerError)
		return
	} 

	var userID, errGetID = GetUserID(application.UserName)
	if errGetID != nil {
		http.Error(w, "Usuario o contraseña incorrectos", http.StatusUnauthorized)
		return
	}

	var salt, errGetSalt = GetSalt(application.UserName)
	if errGetSalt != nil {
		http.Error(w, "Error en el servidor al recuperar el salt", http.StatusInternalServerError)
		return
	}
	
	var storeHash, errGetHash =  GetHashPassword(application.UserName)
	if errGetHash != nil {
		http.Error(w, "Error en el servidor al recuperar el hash", http.StatusInternalServerError)
		return
	}

	var passwordWithSalt = append([]byte(application.Password), salt...)
	var errComparePassword = bcrypt.CompareHashAndPassword([]byte(storeHash), passwordWithSalt)
	if errComparePassword != nil {
		http.Error(w, "Usuario o contraseña incorrectos", http.StatusUnauthorized)
		return
	}

	var token, errToken = services.GenerateJWT(userID, application.UserName)
	if errToken != nil {
		http.Error(w, "Error al generar el token", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Inicio de sesión exitoso",
		"token":   token,
	})
}