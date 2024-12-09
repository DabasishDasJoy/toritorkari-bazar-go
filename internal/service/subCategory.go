package service

import (
	"errors"
	"toritorkari-bazar/internal/domain"
	"toritorkari-bazar/internal/models"
	"toritorkari-bazar/types"

	"gorm.io/gorm"
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

func (service SubCategoryService) GetSubCategory(subCategoryId uint) (types.SubCategoryRequest, error) {
	subCategory, err := service.repo.GetSubCategory(subCategoryId)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return types.SubCategoryRequest{}, errors.New("sub category not found")
		}
	}

	return types.SubCategoryRequest{
		ID:         subCategory.ID,
		Name:       subCategory.Name,
		CategoryId: subCategory.CategoryId,
	}, nil
}
