package domain

import (
	"toritorkari-bazar/internal/models"
	"toritorkari-bazar/types"
)

type IReviewRepo interface {
	CreateReview(review models.Review) (models.Review, error)
	GetReviews(productId uint) ([]models.ReviewResponse, error)
}

type IReviewService interface {
	CreateReview(review types.ReviewRequest) (types.ReviewRequest, error)
	GetReviews(productId uint) ([]types.ReviewResponse, error)
}
