package domain

import (
	"toritorkari-bazar/internal/models"
	"toritorkari-bazar/types"
)

type IProductRepo interface {
	CreateProducts(product []models.Product) error
	GetProducts(categoryId uint) []models.Product
}

type IProductService interface {
	CreateProducts(products []types.ProductRequest) error
	GetProducts(categoryId uint) ([]types.ProductRequest, error)
}
