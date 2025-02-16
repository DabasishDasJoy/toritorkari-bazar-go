package cmd

import (
	"fmt"
	"log"
	connection "toritorkari-bazar/client/conn"
	"toritorkari-bazar/config"
	"toritorkari-bazar/internal/http/controllers"
	"toritorkari-bazar/internal/http/routes"
	"toritorkari-bazar/internal/repositories"
	"toritorkari-bazar/internal/service"

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
	subCategoryController := controllers.SetSubCategoryService(subCategoryService, categoryService)
	routes.SubCategoryRoutes(e, subCategoryController)

	productRepo := repositories.ProductDBInstance(db)
	productService := service.ProductServiceInstance(productRepo)
	productController := controllers.ProductServiceInstance(productService, subCategoryService)
	routes.ProductRouts(e, productController)

	userRepo := repositories.UserDBInstance(db)
	userService := service.UserServiceInstance(userRepo)
	controllers.SetUserServiceInstance(userService)
	routes.UserRoutes(e)

	reviewRepo := repositories.ReviewDBInstance(db)
	reviewService := service.ReviewServiceInstance(reviewRepo)
	reviewController := controllers.SetReviewController(reviewService, productService)
	routes.ReviewRoutes(e, reviewController)

	log.Fatal(e.Start(fmt.Sprintf(":%s", config.LocalConfig.Port)))

}
