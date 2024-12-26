package routes

import (
	"toritorkari-bazar/internal/http/controllers"

	"github.com/labstack/echo/v4"
)

func ProductRouts(e *echo.Echo, productController *controllers.ProductController) {
	Product := e.Group("product")
	Product.POST("", productController.CreateProducts)
	Product.GET("", productController.GetProducts)
}
