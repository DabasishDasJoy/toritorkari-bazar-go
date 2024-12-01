package repositories

import (
	"toritorkari-bazar/pkg/domain"
	"toritorkari-bazar/pkg/models"

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

func (repo *SubCategoryRepo) CreateSubCategories(subCategories []*models.SubCategory) error {
	if err := repo.db.Create(subCategories).Error; err != nil {
		return err
	}

	return nil
}
