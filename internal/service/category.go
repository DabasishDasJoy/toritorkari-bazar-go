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

func (service CategoryService) GetCategories(categoryId uint) ([]types.CategoryResponse, error) {
	var Categories []types.CategoryResponse
	categoryMap := make(map[uint]uint)

	categories := service.repo.GetCategories(categoryId)

	if len(categories) == 0 {
		return nil, errors.New("no categories found")
	}

	for _, val := range categories {
		// Check if category already exists in the map
		if index, exists := categoryMap[val.ID]; exists {

			Categories[index].SubCategories = append(Categories[index].SubCategories, types.SubCategoryRequest{
				ID:         val.SubCategoryID,
				Name:       val.SubCategoryName,
				CategoryId: val.ID,
			})
		} else {
			newCategory := types.CategoryResponse{
				ID:            val.ID,
				Name:          val.Name,
				Icon:          val.Icon,
				SubCategories: make([]types.SubCategoryRequest, 0),
			}

			if val.SubCategoryID != 0 {
				newCategory.SubCategories = append(newCategory.SubCategories, types.SubCategoryRequest{
					ID:         val.SubCategoryID,
					Name:       val.SubCategoryName,
					CategoryId: val.ID,
				})
			}

			Categories = append(Categories, newCategory)
			categoryMap[val.ID] = uint(len(Categories) - 1)
		}
	}

	return Categories, nil
}
