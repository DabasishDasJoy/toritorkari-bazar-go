package service

import (
	"errors"
	"toritorkari-bazar/pkg/domain"
	"toritorkari-bazar/pkg/models"
	"toritorkari-bazar/pkg/types"
)

type CategoryService struct {
	repo domain.ICategoryRepo
}

func CategoryServiceInstance(categoryRepo domain.ICategoryRepo) domain.ICategoryService {
	return &CategoryService{
		repo: categoryRepo,
	}
}

func (service *CategoryService) CreateCategories(categories []*models.Category) error {
	if err := service.repo.CreateCategories(categories); err != nil {
		return errors.New("categories not created")
	}

	return nil
}

func (service *CategoryService) GetCategories() ([]types.CategoryRequest, error) {
	var Categories []types.CategoryRequest

	categories := service.repo.GetCategories()

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
