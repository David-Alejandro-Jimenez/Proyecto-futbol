package router

import (
	"net/http"

	protectedRoutes "github.com/David-Alejandro-Jimenez/Pagina-futbol/internal/api/handlers/private"
	"github.com/David-Alejandro-Jimenez/Pagina-futbol/internal/api/handlers/public"
	"github.com/David-Alejandro-Jimenez/Pagina-futbol/internal/middleware"
	"github.com/gorilla/mux"
)

func SetupRoutes() *mux.Router {
	var router = mux.NewRouter()

	//Uso de JS y CSS
	staticDir := "./../frontend"
	router.PathPrefix("/css/").Handler(http.StripPrefix("/css/", http.FileServer(http.Dir(staticDir+"/css/"))))
	router.PathPrefix("/js/").Handler(http.StripPrefix("/js/", http.FileServer(http.Dir(staticDir+"/js/"))))
	router.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", http.FileServer(http.Dir(staticDir+"/assets/"))))

	//rutas publicas
	router.HandleFunc("/", public.Main_Page).Methods("GET")
	router.HandleFunc("/register", public.RegisterPOST).Methods("POST")
	router.HandleFunc("/login", public.LoginGET).Methods("GET")
	router.HandleFunc("/login", public.LoginPOST).Methods("POST")
	router.HandleFunc("/register", public.RegisterGET).Methods("GET")

	//rutas protegidas
	router.Handle("/home", middleware.AuthMiddleware(http.HandlerFunc(protectedRoutes.HomeHandler))).Methods("GET")
	return router
}