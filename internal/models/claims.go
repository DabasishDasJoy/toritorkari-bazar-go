package models

import "github.com/dgrijalva/jwt-go"

type Claims struct {
	// Role string `json:"role"`
	UserID uint `json:"userID"`
	jwt.StandardClaims
}
