package repositories

import (
	"toritorkari-bazar/internal/domain"
	"toritorkari-bazar/internal/models"

	"gorm.io/gorm"
)

type ProductRepo struct {
	db *gorm.DB
}

func ProductDBInstance(db *gorm.DB) domain.IProductRepo {
	return &ProductRepo{
		db: db,
	}
}

func (repo ProductRepo) CreateProducts(products []models.Product) error {
	if err := repo.db.Create(products).Error; err != nil {
		return err
	}

	return nil
}
