package public

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/David-Alejandro-Jimenez/Pagina-futbol/internal/repository"
	"github.com/David-Alejandro-Jimenez/Pagina-futbol/internal/services"
	"github.com/David-Alejandro-Jimenez/Pagina-futbol/models"
)

func RegisterGET(w http.ResponseWriter, r *http.Request) {
	filePath := "./../frontend/pages/register.html"
	if _, err := os.Stat(filePath); err != nil {
		http.Error(w, "Archivo no encontrado", http.StatusNotFound)
		return
	}
	http.ServeFile(w, r, filePath)
}

func RegisterPOST(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Disallowed method", http.StatusMethodNotAllowed)
		return
	}

	var application models.Account
	err = json.NewDecoder(r.Body).Decode(&application)
	if err != nil {
		http.Error(w, "Invalid Request", http.StatusBadRequest)
		return
	}

	err = services.ValidateUserName(application.UserName)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	
	err = services.ValidatePassword(application.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	exists, err := repository.GetUser(application.UserName) 
	if err != nil {
    	http.Error(w, "Server error while validating user", http.StatusInternalServerError)
   		return
	}

	if exists {
		http.Error(w, "Username already exists", http.StatusConflict)
		return
	}

	err = repository.SaveUser(application.UserName, application.Password)
	if err != nil {
		http.Error(w, "Error saving user in database", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "User created successfully"})
}