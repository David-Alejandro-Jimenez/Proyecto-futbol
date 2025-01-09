package public

import (
	"encoding/json"
	"net/http"
	"os"
	"time"

	"github.com/David-Alejandro-Jimenez/Pagina-futbol/internal/repository"
	"github.com/David-Alejandro-Jimenez/Pagina-futbol/internal/services"
	"github.com/David-Alejandro-Jimenez/Pagina-futbol/models"
	"golang.org/x/crypto/bcrypt"
)

var err error

func LoginGET(w http.ResponseWriter, r *http.Request) {
	filePath := "./../frontend/pages/login.html"
	if _, err := os.Stat(filePath); err != nil {
		http.Error(w, "Archivo no encontrado", http.StatusNotFound)
		return
	}
	http.ServeFile(w, r, filePath)
}

func LoginPOST(w http.ResponseWriter, r *http.Request) {
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

	err = services.ValidateUserName(application.UserName)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	exists, err := repository.GetUser(application.UserName)
	if err != nil {
		http.Error(w, "Error en el servidor al validar el usuario", http.StatusInternalServerError)
		return
		} 
		
	if !exists {
		http.Error(w, "Usuario o contraseña incorrectos", http.StatusUnauthorized)
		return
	}

	salt, err := repository.GetSalt(application.UserName)
	if err != nil {
		http.Error(w, "Error en el servidor al recuperar el salt", http.StatusInternalServerError)
		return
	}
	
	storeHash, err :=  repository.GetHashPassword(application.UserName)
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