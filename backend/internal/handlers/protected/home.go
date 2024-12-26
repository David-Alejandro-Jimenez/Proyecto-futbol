package protected

import (
	"fmt"
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
		http.Error(w, "cookie not found or invalid", http.StatusUnauthorized)
		return
	}

	token, err :=  jwt.ParseWithClaims(cookie.Value, &models.Claims{}, func(t *jwt.Token) (interface{}, error) {
		_, ok := t.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, fmt.Errorf("invalid signature method")
		}

		return jwtSecret, nil 
	})

	if err != nil || !token.Valid {
		http.Error(w, "Unauthorized access", http.StatusUnauthorized,)
		return
	}

	claims, ok := token.Claims.(*models.Claims)
	if !ok {
		http.Error(w, "Invalid claims structure", http.StatusUnauthorized)
    	return
	}

	w.Write([]byte(fmt.Sprintf("¡Hola usuario %v! B", claims.UserName)))
}