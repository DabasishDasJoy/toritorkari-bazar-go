package service

import (
	"errors"
	"toritorkari-bazar/internal/domain"
	"toritorkari-bazar/internal/models"
	"toritorkari-bazar/types"
)

type SubCategoryService struct {
	repo domain.ISubCategoryRepo
}

func SubCategoryServiceInstance(subCategoryRepo domain.ISubCategoryRepo) domain.ISubCategoryService {
	return &SubCategoryService{
		repo: subCategoryRepo,
	}
}

func (service SubCategoryService) CreateSubCategories(reqSubCategory []types.SubCategoryRequest) error {
	subCategories := make([]models.SubCategory, 0, len(reqSubCategory))

	for _, subCategory := range reqSubCategory {
		subCategories = append(subCategories, models.SubCategory{
			Name:       subCategory.Name,
			CategoryId: subCategory.CategoryId,
		})
	}

	if err := service.repo.CreateSubCategories(subCategories); err != nil {
		return errors.New("sub categories not created")
	}

	return nil
}
