package main

import (
	"log"
	"net/http"

	"github.com/David-Alejandro-Jimenez/Pagina-futbol/internal/config"
	"github.com/David-Alejandro-Jimenez/Pagina-futbol/internal/repository/database"
	"github.com/David-Alejandro-Jimenez/Pagina-futbol/internal/api"
)

func main() {
	var errConfig = config.LoadConfig()
	if errConfig != nil {
		log.Fatalf("Error al cargar la configuración: %v", errConfig)
	}

	//Inicio de base de datos
	var errdb = database.InitDB()
	if errdb != nil {
		log.Println("No se conecto a la base de datos")
	}
	defer database.DB.Close()

	var router = router.SetupRoutes()
	
	//Iniciar Servidor
	var port = ":8080"
	log.Printf("Servidor escuchando en http://localhost%s", port)
	var err = http.ListenAndServe(port, router)
	if err != nil {
		log.Fatalf("Error al iniciar el servidor: %v", err)
	}
}