package routes

import (
	"toritorkari-bazar/internal/http/controllers"

	"github.com/labstack/echo/v4"
)

func SubCategoryRoutes(e *echo.Echo) {
	subCategory := e.Group("/sub-category")

	subCategory.POST("", controllers.CreateSubCategories)
}
