package containers

import (
	"fmt"
	"log"
	"toritorkari-bazar/pkg/config"
	"toritorkari-bazar/pkg/connection"
	"toritorkari-bazar/pkg/controllers"
	"toritorkari-bazar/pkg/repositories"
	"toritorkari-bazar/pkg/routes"
	"toritorkari-bazar/pkg/service"

	"github.com/labstack/echo/v4"
)

func Serve(e *echo.Echo) {
	config.SetConfig()

	db := connection.GetDB()

	bookRepo := repositories.BookDBInstance(db)
	bookService := service.BookServiceInstance(bookRepo)
	controllers.SetBookService(bookService)
	routes.BookRoutes(e)

	categoryRepo := repositories.CategoryDBInstance(db)
	categoryService := service.CategoryServiceInstance(categoryRepo)
	controllers.SetCategoryService(categoryService)
	routes.CategoryRoutes(e)

	subCategoryRepo := repositories.SubCategoryDBInstance(db)
	subCategoryService := service.SubCategoryServiceInstance(subCategoryRepo)
	controllers.SetSubCategoryService(subCategoryService)
	routes.SubCategoryRoutes(e)

	log.Fatal(e.Start(fmt.Sprintf(":%s", config.LocalConfig.Port)))

}
