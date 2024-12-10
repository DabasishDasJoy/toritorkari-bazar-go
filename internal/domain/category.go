package domain

import (
	"toritorkari-bazar/internal/models"
	"toritorkari-bazar/types"
)

type ICategoryRepo interface {
	CreateCategories(categories []models.Category) error
	GetCategories(categoryId uint) []models.IntermediateCategoryResponse
}

type ICategoryService interface {
	CreateCategories(categories []types.CategoryRequest) error
	GetCategories(categoryId uint) ([]types.CategoryResponse, error)
}
