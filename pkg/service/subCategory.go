package service

import (
	"errors"
	"toritorkari-bazar/pkg/domain"
	"toritorkari-bazar/pkg/models"
)

type SubCategoryService struct {
	repo domain.ISubCategoryRepo
}

func SubCategoryServiceInstance(subCategoryRepo domain.ISubCategoryRepo) domain.ISubCategoryRepo {
	return &SubCategoryService{
		repo: subCategoryRepo,
	}
}

func (service *SubCategoryService) CreateSubCategories(subCategories []*models.SubCategory) error {
	if err := service.repo.CreateSubCategories(subCategories); err != nil {
		return errors.New("sub categories not created")
	}

	return nil
}
