package protected

import (
	"fmt"
	"log"
	"net/http"

	"github.com/David-Alejandro-Jimenez/Pagina-futbol/models"
	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/viper"
)

var jwtSecret = []byte(viper.GetString("JWT_SECRET_KEY"))

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Entrando al HomeHandler")
	for name, values := range r.Header {
		fmt.Printf("%s: %s\n", name, values)
	}
	

	cookie, err := r.Cookie("token")
	if err != nil {
		log.Println("Cookie:", cookie.Value)
		http.Error(w, "No se encontró la cookie o no es válida", http.StatusUnauthorized)
		return
	}
	log.Println("Cookie:", cookie.Value)

	token, err :=  jwt.ParseWithClaims(cookie.Value, &models.Claims{}, func(t *jwt.Token) (interface{}, error) {
		_, ok := t.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, fmt.Errorf("método de firma inválido")
		}

		return jwtSecret, nil 
	})
	log.Println("Token:", token.Valid, err)

	if err != nil || !token.Valid {
		log.Println("Token:", token.Valid, err)
		http.Error(w, "Acceso no autorizado", http.StatusUnauthorized,)
		return
	}

	claims, ok := token.Claims.(*models.Claims)
	if !ok {
		http.Error(w, "Estructura de claims inválida", http.StatusUnauthorized)
    	return
	}
	log.Println("Claims:", claims.UserName)

	w.Write([]byte(fmt.Sprintf("¡Hola usuario %v! B", claims.UserName)))
}