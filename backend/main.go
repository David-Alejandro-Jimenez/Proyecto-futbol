package main

import (
	"log"
	"net/http"

	db "github.com/David-Alejandro-Jimenez/Pagina-futbol/database"
	"github.com/David-Alejandro-Jimenez/Pagina-futbol/handlers"
	"github.com/David-Alejandro-Jimenez/Pagina-futbol/internal/config"
	"github.com/gorilla/mux"
)

func main() {
	var errConfig = config.LoadConfig()
	if errConfig != nil {
		log.Fatalf("Error al cargar la configuración: %v", errConfig)
	}

	//Inicio de base de datos
	var errdb = db.InitDB()
	if errdb != nil {
		log.Println("No se conecto a la base de datos")
	}
	defer db.DataBase.Close()

	//Rutas
	var router = mux.NewRouter()
	router.HandleFunc("/createAccount", handlers.RegisterNewAccount).Methods("POST")

	//Iniciar Servidor
	var port = ":8080"
	log.Printf("Servidor escuchando en http://localhost%s", port)
	var err = http.ListenAndServe(port, router)
	if err != nil {
		log.Fatalf("Error al iniciar el servidor: %v", err)
	}
}