package middleware

import (
	"net/http"
	"strings"
	"toritorkari-bazar/config"
	"toritorkari-bazar/internal/models"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

// AuthenticateUser middleware validates the JWT token and ensures the user is authenticated
func AuthenticateUser() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			configs := config.LocalConfig

			// Retrieve the Authorization header
			authHeader := c.Request().Header.Get("Authorization")
			if authHeader == "" {
				return echo.NewHTTPError(http.StatusUnauthorized, "unauthorized")
			}

			// Ensure the token uses the Bearer scheme
			parts := strings.Split(authHeader, " ")
			if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
				return echo.NewHTTPError(http.StatusUnauthorized, "unauthorized")
			}

			// Parse the token
			tokenString := parts[1]
			claims := &models.Claims{}
			token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
				// Ensure the signing method is correct
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, echo.NewHTTPError(http.StatusUnauthorized, "unauthorized")
				}
				return []byte(configs.JWTKEY), nil
			})

			// Handle token validation errors
			if err != nil || !token.Valid {
				return echo.NewHTTPError(http.StatusUnauthorized, "unauthorized")
			}

			// Attach the claims to the context
			c.Set("user", claims)

			// Proceed to the next handler
			return next(c)
		}
	}
}
