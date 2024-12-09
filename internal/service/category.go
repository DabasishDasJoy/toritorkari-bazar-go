package service

import (
	"errors"

	"toritorkari-bazar/internal/domain"
	"toritorkari-bazar/internal/models"
	"toritorkari-bazar/types"
)

type CategoryService struct {
	repo domain.ICategoryRepo
}

func CategoryServiceInstance(categoryRepo domain.ICategoryRepo) domain.ICategoryService {
	return &CategoryService{
		repo: categoryRepo,
	}
}

func (service CategoryService) CreateCategories(reqCategories []types.CategoryRequest) error {
	categories := make([]models.Category, 0, len(reqCategories))
	for _, category := range reqCategories {
		categories = append(categories, models.Category{
			Name: category.Name,
			Icon: category.Icon,
		})
	}

	if err := service.repo.CreateCategories(categories); err != nil {
		return errors.New("categories not created")
	}

	return nil
}

func (service CategoryService) GetCategories(categoryId uint) ([]types.CategoryRequest, error) {
	var Categories []types.CategoryRequest

	categories := service.repo.GetCategories(categoryId)

	if len(categories) == 0 {
		return nil, errors.New("no categories found")
	}

	for _, val := range categories {
		Categories = append(Categories, types.CategoryRequest{
			ID:   val.ID,
			Name: val.Name,
			Icon: val.Icon,
		})
	}

	return Categories, nil
}
