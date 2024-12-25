package auth

import (
	"log"
	"net/http"

	"github.com/David-Alejandro-Jimenez/Pagina-futbol/services"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("token")
		if err != nil || cookie.Value == "" {
			http.Error(w, "No autorizado", http.StatusForbidden)
			return
		}

		tokenString := cookie.Value
		log.Println(tokenString)
		_, err = services.ValidateToken(tokenString) 
		if err != nil {
			log.Println(err)
			http.Error(w, "Token inválido o expirado", http.StatusForbidden)
			return
		}

	next.ServeHTTP(w, r)
	})
}