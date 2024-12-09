package domain

import (
	"toritorkari-bazar/internal/models"
	"toritorkari-bazar/types"
)

type ISubCategoryRepo interface {
	CreateSubCategories(subCategory []models.SubCategory) error
}

type ISubCategoryService interface {
	CreateSubCategories(subCategory []types.SubCategoryRequest) error
}