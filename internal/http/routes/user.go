package routes

import (
	"toritorkari-bazar/internal/http/controllers"

	"github.com/labstack/echo/v4"
)

func UserRoutes(e *echo.Echo) {
	user := e.Group("/user")

	user.POST("/signup", controllers.SignUp)
	user.POST("/login", controllers.Login)
}
