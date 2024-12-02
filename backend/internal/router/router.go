package router

import (
	"github.com/David-Alejandro-Jimenez/Pagina-futbol/internal/auth"
	"github.com/gorilla/mux"
)

func SetupRoutes() *mux.Router {
	var router = mux.NewRouter()
	router.HandleFunc("/register", auth.RegisterNewAccount).Methods("POST")
	router.HandleFunc("/login", auth.LoginIn).Methods("POST")

	return router
}