package controllers

import (
	"net/http"
	"strconv"
	"toritorkari-bazar/internal/domain"
	"toritorkari-bazar/internal/models"
	"toritorkari-bazar/types"

	"github.com/labstack/echo/v4"
)

type ReviewController struct {
	ReviewService  domain.IReviewService
	ProductService domain.IProductService
}

func SetReviewController(reviewService domain.IReviewService, productService domain.IProductService) *ReviewController {
	return &ReviewController{
		ReviewService:  reviewService,
		ProductService: productService,
	}
}

func (reviewController ReviewController) CreateReview(e echo.Context) error {
	user, ok := e.Get("user").(*models.Claims)

	if !ok {
		return e.JSON(http.StatusUnauthorized, map[string]string{
			"error": "Unauthorized",
		})
	}

	requestReview := types.ReviewRequest{}

	if err := e.Bind(&requestReview); err != nil {
		return e.JSON(http.StatusBadRequest, "invalid data")
	}

	// validate product
	if _, err := reviewController.ProductService.GetProduct(requestReview.ProductID); err != nil {
		return e.JSON(http.StatusBadRequest, "invalid product id")
	}

	// validate review request
	if err := requestReview.ValidateReview(); err != nil {
		return e.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	requestReview.UserID = user.UserID

	review, err := reviewController.ReviewService.CreateReview(requestReview)

	if err != nil {
		return e.JSON(http.StatusInternalServerError, err.Error())
	}

	return e.JSON(http.StatusCreated, review)
}

func (reviewController ReviewController) GetReviews(e echo.Context) error {
	temporaryProductId := e.Param("productID")

	productID, err := strconv.ParseInt(temporaryProductId, 0, 0)

	if err != nil {
		return e.JSON(http.StatusBadRequest, "invalid product id")
	}

	reviews, err := reviewController.ReviewService.GetReviews(uint(productID))
	if err != nil {
		return e.JSON(http.StatusInternalServerError, err.Error())
	}

	return e.JSON(http.StatusFound, reviews)
}
