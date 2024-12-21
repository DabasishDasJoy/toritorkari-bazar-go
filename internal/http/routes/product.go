package routes

import (
	"toritorkari-bazar/internal/http/controllers"
	middleware "toritorkari-bazar/internal/http/middlewares"

	"github.com/labstack/echo/v4"
)

func ProductRouts(e *echo.Echo) {
	Product := e.Group("product")
	Product.Use(middleware.AuthenticateUser())
	Product.POST("", controllers.CreateProducts)
	Product.GET("/list", controllers.GetProducts)
}
