package routes

import (
	"toritorkari-bazar/internal/http/controllers"
	middleware "toritorkari-bazar/internal/http/middlewares"

	"github.com/labstack/echo/v4"
)

func ProductRouts(e *echo.Echo, productController *controllers.ProductController) {
	Product := e.Group("product")
	Product.Use(middleware.AuthenticateUser())
	Product.POST("", productController.CreateProducts)
	Product.GET("/list", productController.GetProducts)
}
