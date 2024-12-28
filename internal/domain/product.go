package domain

import (
	"toritorkari-bazar/internal/models"
	"toritorkari-bazar/types"
)

type IProductRepo interface {
	CreateProducts(product []models.Product) error
	GetProducts(types.GetCategoriesParams) ([]models.Product, int, error)
	GetProduct(id uint) (models.Product, error)
}

type IProductService interface {
	CreateProducts(products []types.ProductRequest) error
	GetProducts(types.GetCategoriesParams) (types.ProductResponse, error)
	GetProduct(id uint) (types.ProductRequest, error)
}
