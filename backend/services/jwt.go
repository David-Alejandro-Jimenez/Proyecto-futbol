package services

import (
	"time"

	"github.com/David-Alejandro-Jimenez/Pagina-futbol/models"
	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/viper"
)

var jwtSecret = []byte(viper.GetString("JWT_SECRET_KEY"))

func GenerateJWT(userID int, userName string) (string, error) {
	var claims = models.Claims{
		UserID: userID,
		UserName: userName,
			RegisteredClaims: jwt.RegisteredClaims{
				ExpiresAt: jwt.NewNumericDate(time.Now().Add(12 * time.Hour)),
			},
	}

	var token = jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

func ValidateToken(tokenString string) (*models.Claims, error) {
	var token, err = jwt.ParseWithClaims(tokenString, models.Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if err != nil || !token.Valid {
		return nil, err
	}

	var claims, ok = token.Claims.(*models.Claims)
	if !ok {
		return nil, err
	}

	return claims, nil
}