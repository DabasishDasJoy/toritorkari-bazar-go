package repositories

import (
	"toritorkari-bazar/pkg/domain"
	"toritorkari-bazar/pkg/models"

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

func (repo *CategoryRepo) CreateCategories(categories []*models.Category) error {
	if err := repo.db.Create(categories).Error; err != nil {
		return err
	}

	return nil
}

func (repo *CategoryRepo) GetCategories() []models.Category {
	var Categories []models.Category

	err := repo.db.Find(&Categories).Error

	if err != nil {
		return []models.Category{}
	}

	return Categories
}
