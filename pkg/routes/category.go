package routes

import (
	"toritorkari-bazar/pkg/controllers"

	"github.com/labstack/echo/v4"
)

func CategoryRoutes(e *echo.Echo) {
	category := e.Group("/category")

	category.POST("", controllers.CreateCategories)
	category.GET("", controllers.GetCategories)
}
