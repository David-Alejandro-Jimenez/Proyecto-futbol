package router

import (
	"net/http"

	"github.com/David-Alejandro-Jimenez/Pagina-futbol/internal/auth"
	protectedRoutes "github.com/David-Alejandro-Jimenez/Pagina-futbol/internal/handlers/protected"
	"github.com/David-Alejandro-Jimenez/Pagina-futbol/internal/handlers/public"
	"github.com/gorilla/mux"
)

func SetupRoutes() *mux.Router {
	var router = mux.NewRouter()

	//Uso de JS y CSS
	staticDir := "./../frontend"
	router.PathPrefix("/CSS/").Handler(http.StripPrefix("/CSS/", http.FileServer(http.Dir(staticDir+"/CSS/"))))
	router.PathPrefix("/JS/").Handler(http.StripPrefix("/JS/", http.FileServer(http.Dir(staticDir+"/JS/"))))
	router.PathPrefix("/Images/").Handler(http.StripPrefix("/Images/", http.FileServer(http.Dir(staticDir+"/Images/"))))

	//rutas publicas
	router.HandleFunc("/", public.Main_Page).Methods("GET")
	router.HandleFunc("/register", auth.RegisterNewAccount).Methods("POST")
	router.HandleFunc("/login", auth.LoginInGET).Methods("GET")
	router.HandleFunc("/login", auth.LoginInPOST).Methods("POST")

	//rutas protegidas
	var protected = router.PathPrefix("/home").Subrouter()
	protected.HandleFunc("", protectedRoutes.HomeHandler).Methods("GET")
	protected.Use(auth.AuthMiddleware)

	return router
}