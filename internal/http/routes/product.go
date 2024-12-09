package routes

import (
	"toritorkari-bazar/internal/http/controllers"

	"github.com/labstack/echo/v4"
)

func ProductRouts(e *echo.Echo) {
	Product := e.Group("product")

	Product.POST("", controllers.CreateProducts)
}
