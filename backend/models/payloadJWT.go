package models

import "github.com/golang-jwt/jwt/v5"

type Claims struct {
	UserID   int    `json:"UserID"`
	UserName string `json:"UserName"`
	jwt.RegisteredClaims
}