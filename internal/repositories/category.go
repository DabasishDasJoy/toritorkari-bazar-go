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

func (repo CategoryRepo) GetCategories(categoryId uint) []models.Category {
	var Categories []models.Category
	var err error

	if categoryId != 0 {
		err = repo.db.Where("id = ?", categoryId).Find(&Categories).Error
	} else {
		err = repo.db.Find(&Categories).Error
	}

	if err != nil {
		return []models.Category{}
	}

	return Categories
}
