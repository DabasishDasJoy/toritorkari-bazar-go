package domain

import (
	"toritorkari-bazar/internal/models"
	"toritorkari-bazar/types"
)

type ICategoryRepo interface {
	CreateCategories(categories []models.Category) error
	GetCategories(categoryId uint) []models.Category
}

type ICategoryService interface {
	CreateCategories(categories []types.CategoryRequest) error
	GetCategories(categoryId uint) ([]types.CategoryRequest, error)
}
