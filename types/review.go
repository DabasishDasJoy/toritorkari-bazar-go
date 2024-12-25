package types

import (
	"time"

	"github.com/go-ozzo/ozzo-validation/is"
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type ReviewRequest struct {
	ID        uint    `json:"id"`
	Rating    float64 `json:"rating"`
	Review    string  `json:"review"`
	ProductID uint    `json:"productID"`
	UserID    uint    `json:"UserID"`
}

type ReviewResponse struct {
	ID        uint    `json:"id"`
	Rating    float64 `json:"rating"`
	Review    string  `json:"review"`
	ProductID uint    `json:"productId"`
	CreatedAt time.Time
	User      UserResponse
}

func (review ReviewRequest) ValidateReview() error {
	return validation.ValidateStruct(&review,
		validation.Field(&review.ProductID,
			validation.Required.Error("ProductID cannot be empty"),
		),
		validation.Field(&review.Rating,
			validation.Required.Error("Rating cannot be empty"),
			validation.Max(5.0).Error("rating can be max 5"),
		),

		validation.Field(&review.Review,
			is.Alphanumeric.Error("Review must be string"),
		),
	)
}
