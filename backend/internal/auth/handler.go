package auth

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/David-Alejandro-Jimenez/Pagina-futbol/models"
	"github.com/David-Alejandro-Jimenez/Pagina-futbol/services"
	"golang.org/x/crypto/bcrypt"
)

var err error

func RegisterNewAccount(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Disallowed method", http.StatusMethodNotAllowed)
		return
	}

	var application models.Account
	err = json.NewDecoder(r.Body).Decode(&application)
	if err != nil {
		log.Println(err)
		http.Error(w, "Invalid Request", http.StatusBadRequest)
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

	exists, err := GetUser(application.UserName) 
	if err != nil {
    	http.Error(w, "Server error while validating user", http.StatusInternalServerError)
   		return
	}

	if exists {
		http.Error(w, "Username already exists", http.StatusConflict)
		return
	}

	err = SaveUser(application.UserName, application.Password)
	if err != nil {
		http.Error(w, "Error saving user in database", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "User created successfully"})
}

func LoginInGET(w http.ResponseWriter, r *http.Request) {
	filePath := "./../frontend/RegistroUsuarios/LoginIn.html"
	if _, err := os.Stat(filePath); err != nil {
		http.Error(w, "Archivo no encontrado", http.StatusNotFound)
		return
	}
	http.ServeFile(w, r, filePath)
}

func LoginInPOST(w http.ResponseWriter, r *http.Request) {
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

	exists, err := GetUser(application.UserName)
	if err != nil {
		http.Error(w, "Error en el servidor al validar el usuario", http.StatusInternalServerError)
		return
		} 
		
	if !exists {
		http.Error(w, "Usuario o contraseña incorrectos", http.StatusUnauthorized)
		return
	}

	salt, err := GetSalt(application.UserName)
	if err != nil {
		http.Error(w, "Error en el servidor al recuperar el salt", http.StatusInternalServerError)
		return
	}
	
	storeHash, err :=  GetHashPassword(application.UserName)
	if err != nil {
		http.Error(w, "Error en el servidor al recuperar el hash", http.StatusInternalServerError)
		return
	}

	var passwordWithSalt = append([]byte(application.Password), salt...)
	err = bcrypt.CompareHashAndPassword([]byte(storeHash), passwordWithSalt)
	if err != nil {
		http.Error(w, "Usuario o contraseña incorrectos", http.StatusUnauthorized)
		return
	}

	token, err := services.GenerateJWT(application.UserName)
	if err != nil {
		http.Error(w, "Error al generar el token", http.StatusInternalServerError)
		return
	}

	cookie := http.Cookie{
		Name: "token",
		Value: token,
		Expires: time.Now().Add(12 * time.Hour),
		HttpOnly: false,
		Path: "/",
		Secure: false,
		SameSite: http.SameSiteLaxMode,
	}

	http.SetCookie(w, &cookie) 

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Inicio de sesión exitoso",
		"token":   token,
		"redirect": "/home",
	})
}