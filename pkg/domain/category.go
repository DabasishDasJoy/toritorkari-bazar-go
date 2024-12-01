package domain

import (
	"toritorkari-bazar/pkg/models"
	"toritorkari-bazar/pkg/types"
)

type ICategoryRepo interface {
	CreateCategories(categories []*models.Category) error
	GetCategories() []models.Category
}

type ICategoryService interface {
	CreateCategories(categories []*models.Category) error
	GetCategories() ([]types.CategoryRequest, error)
}
