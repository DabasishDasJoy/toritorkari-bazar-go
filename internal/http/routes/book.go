package routes

import (
	"toritorkari-bazar/internal/http/controllers"

	"github.com/labstack/echo/v4"
)

func BookRoutes(e *echo.Echo) {
	book := e.Group("/bookstore")

	book.POST("/book", controllers.CreateBook)
	book.GET("/book", controllers.GetBooks)
	// book.PUT("/book/:bookId", controllers.UpdateBook)
	// book.DELETE("/book/:bookId", controllers.DeleteBook)
}
