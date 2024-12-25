package methods

import (
	"errors"
	"time"
	"toritorkari-bazar/config"
	"toritorkari-bazar/internal/models"

	"github.com/dgrijalva/jwt-go"
)

func GenerateToken(email string, expiration time.Duration, userId uint) (string, error) {
	configs := config.LocalConfig
	expirationTime := time.Now().Add(expiration)

	claims := models.Claims{
		UserID: userId,
		StandardClaims: jwt.StandardClaims{
			Subject:   email,
			ExpiresAt: expirationTime.Unix(),
		},
	}

	secretKey := []byte(configs.JWTKEY)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secretKey)
}

func ValidateRefreshToken(tokenString string) (models.Claims, error) {
	configs := config.LocalConfig
	token, err := jwt.ParseWithClaims(tokenString, &models.Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return models.Claims{}, errors.New("invalid token signing method")
		}
		return []byte(configs.JWTKEY), nil
	})

	if err != nil || !token.Valid {
		return models.Claims{}, errors.New("invalid or expired refresh token")
	}

	claims, ok := token.Claims.(*models.Claims)
	if !ok {
		return models.Claims{}, errors.New("could not parse token claims")
	}

	return *claims, nil
}
