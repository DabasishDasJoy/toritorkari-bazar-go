package repositories

import (
	"toritorkari-bazar/internal/domain"
	"toritorkari-bazar/internal/models"

	"gorm.io/gorm"
)

type CategoryRepo struct {
	db *gorm.DB
}

func CategoryDBInstance(db *gorm.DB) domain.ICategoryRepo {
	return &CategoryRepo{
		db: db,
	}
}

func (repo CategoryRepo) CreateCategories(categories []models.Category) error {
	if err := repo.db.Create(categories).Error; err != nil {
		return err
	}

	return nil
}

func (repo CategoryRepo) GetCategories(categoryId uint) []models.IntermediateCategoryResponse {
	var Categories []models.IntermediateCategoryResponse
	var err error

	// query := repo.db.Model(&models.Category{}).
	// 	Select("categories.id, categories.name, categories.icon, sub_categories.id sub_category_id, sub_categories.name sub_category_name").
	// 	Joins("Left Join sub_categories  ON sub_categories.id = categories.id")

	// if categoryId != 0 {
	// 	query = query.Where("categories.id = ?", categoryId)

	// }

	// err = query.Find(&Categories).Error

	// if err != nil {
	// 	return []models.IntermediateCategoryResponse{}
	// }

	query := `
		select categories.id, categories.name, categories.icon, sub_categories.id sub_category_id, sub_categories.name sub_category_name
		from categories
		left join sub_categories ON sub_categories.category_id = categories.id
	`
	var params []interface{}

	// Add condition if categoryId is not 0
	if categoryId != 0 {
		query += " WHERE categories.id = ?"
		params = append(params, categoryId)
	}

	// Execute query with parameters
	err = repo.db.Raw(query, params...).Find(&Categories).Error

	if err != nil {
		return []models.IntermediateCategoryResponse{}
	}

	return Categories
}
