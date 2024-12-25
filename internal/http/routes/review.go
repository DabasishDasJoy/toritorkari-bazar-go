package routes

import (
	"toritorkari-bazar/internal/http/controllers"
	middleware "toritorkari-bazar/internal/http/middlewares"

	"github.com/labstack/echo/v4"
)

func ReviewRoutes(e *echo.Echo, reviewController *controllers.ReviewController) {
	Review := e.Group("review")
	Review.Use(middleware.AuthenticateUser())

	Review.POST("", reviewController.CreateReview)
	Review.GET("/:productID", reviewController.GetReviews)
}
