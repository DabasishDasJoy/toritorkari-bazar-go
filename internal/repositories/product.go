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

func (repo ProductRepo) GetProducts(categoryId uint) []models.Product {
	var Products []models.Product

	query := `select products.id, products.name, products.description, products.category_id, products.sub_category_id, products.icon, products.price, products.quantity, products.discount, products.status from products`

	var params []interface{}

	if categoryId != 0 {
		query += " WHERE products.category_id =?"
		params = append(params, categoryId)
	}

	if err := repo.db.Raw(query, params...).Find(&Products).Error; err != nil {
		return []models.Product{}
	}

	return Products
}
