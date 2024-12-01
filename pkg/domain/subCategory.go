package domain

import "toritorkari-bazar/pkg/models"

type ISubCategoryRepo interface {
	CreateSubCategories(subCategory []*models.SubCategory) error
}

type ISubCategoryService interface {
	CreateSubCategories(subCategory []*models.SubCategory) error
}
