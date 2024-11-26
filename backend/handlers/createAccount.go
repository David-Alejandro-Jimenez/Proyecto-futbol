package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/David-Alejandro-Jimenez/Pagina-futbol/database"
	models "github.com/David-Alejandro-Jimenez/Pagina-futbol/models"
	services "github.com/David-Alejandro-Jimenez/Pagina-futbol/services"
)


func RegisterNewAccount(w http.ResponseWriter, r *http.Request) {
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

	var errValidateEntries = services.ValidateUserName(application.UserName)
	if errValidateEntries != nil {
		http.Error(w, errValidateEntries.Error(), http.StatusBadRequest)
		return
	}
	
	errValidateEntries = services.ValidatePassword(application.Password)
	if errValidateEntries != nil {
		http.Error(w, errValidateEntries.Error(), http.StatusBadRequest)
		return
	}

	var exists, errValidateExistingUsers = database.ValidateExistingUsers(application.UserName) 
	if errValidateExistingUsers != nil {
		log.Printf("Error verificando usuario existente: %v", errValidateExistingUsers)
    	http.Error(w, "Error en el servidor al validar el usuario", http.StatusInternalServerError)
   		return
	}

	if exists {
		http.Error(w, "El nombre de usuario ya existe", http.StatusConflict)
		return
	}

	var errSaveUsers = database.SaveUser(application.UserName, application.Password)
	if errSaveUsers != nil {
		http.Error(w, "Error al guardar el usuario en la base de datos", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Usuario creado exitosamente"})
}
