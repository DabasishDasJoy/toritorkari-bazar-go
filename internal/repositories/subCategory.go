package repositories

import (
	"toritorkari-bazar/internal/domain"
	"toritorkari-bazar/internal/models"

	"gorm.io/gorm"
)

type SubCategoryRepo struct {
	db *gorm.DB
}

func SubCategoryDBInstance(db *gorm.DB) domain.ISubCategoryRepo {
	return &SubCategoryRepo{
		db: db,
	}
}

func (repo SubCategoryRepo) CreateSubCategories(subCategories []models.SubCategory) error {
	if err := repo.db.Create(subCategories).Error; err != nil {
		return err
	}

	return nil
}

func (repo SubCategoryRepo) GetSubCategory(subCategoryId uint) (models.SubCategory, error) {
	var SubCategory models.SubCategory

	err := repo.db.Where("id=?", subCategoryId).First(&SubCategory).Error

	if err != nil {
		return models.SubCategory{}, err
	}

	return SubCategory, nil
}
