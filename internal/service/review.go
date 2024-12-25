package service

import (
	"errors"
	"toritorkari-bazar/internal/domain"
	"toritorkari-bazar/internal/models"
	"toritorkari-bazar/types"
)

type ReviewService struct {
	repo domain.IReviewRepo
}

func ReviewServiceInstance(reviewRepo domain.IReviewRepo) domain.IReviewService {
	return &ReviewService{
		repo: reviewRepo,
	}
}

func (ReviewService ReviewService) CreateReview(review types.ReviewRequest) (types.ReviewRequest, error) {
	reviewRequest := models.Review{
		ProductID: review.ProductID,
		UserID:    review.UserID,
		Rating:    review.Rating,
		Review:    review.Review,
	}

	var (
		newReview models.Review
		err       error
	)

	newReview, err = ReviewService.repo.CreateReview(reviewRequest)

	if err != nil {
		return types.ReviewRequest{}, err
	}

	return types.ReviewRequest{
		ID:        newReview.ID,
		Rating:    newReview.Rating,
		Review:    newReview.Review,
		ProductID: newReview.ProductID,
		UserID:    newReview.UserID,
	}, nil
}

func (ReviewService ReviewService) GetReviews(productId uint) ([]types.ReviewResponse, error) {
	var reviewResponses []types.ReviewResponse

	reviews, err := ReviewService.repo.GetReviews(productId)

	if err != nil {
		return nil, err
	}

	if len(reviews) == 0 {
		return nil, errors.New("no reviews found")
	}

	for _, review := range reviews {
		reviewResponses = append(reviewResponses, types.ReviewResponse{
			ID:        review.Review.ID,
			Rating:    review.Review.Rating,
			Review:    review.Review.Review,
			ProductID: review.ProductID,
			CreatedAt: review.CreatedAt,
			User: types.UserResponse{
				ID:    review.UserID,
				Name:  review.Name,
				Email: review.Email,
			},
		})
	}

	return reviewResponses, nil
}
