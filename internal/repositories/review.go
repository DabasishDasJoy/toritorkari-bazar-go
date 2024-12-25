package repositories

import (
	"toritorkari-bazar/internal/domain"
	"toritorkari-bazar/internal/models"

	"gorm.io/gorm"
)

type ReviewRepo struct {
	db *gorm.DB
}

func ReviewDBInstance(db *gorm.DB) domain.IReviewRepo {
	return &ReviewRepo{
		db: db,
	}
}

func (reviewRepo ReviewRepo) CreateReview(reviewRequest models.Review) (models.Review, error) {
	if err := reviewRepo.db.Create(&reviewRequest).Error; err != nil {
		return models.Review{}, err
	}

	return reviewRequest, nil
}

func (reviewRepo ReviewRepo) GetReviews(productId uint) ([]models.ReviewResponse, error) {
	var revieResponses []models.ReviewResponse

	query := `select reviews.id, reviews.rating, reviews.review, reviews.product_id, reviews.user_id, reviews.created_at, users.id, users.email, users.name from reviews LEFT join users on reviews.user_id=users.id where reviews.product_id=? order by created_at DESC`

	var params interface{}

	if productId != 0 {
		params = productId
	}

	err := reviewRepo.db.Raw(query, params).Find(&revieResponses).Error

	if err != nil {
		return []models.ReviewResponse{}, err
	}

	return revieResponses, nil
}
