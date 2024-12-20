package domain

import (
	"toritorkari-bazar/internal/models"
	"toritorkari-bazar/types"
)

type IProductRepo interface {
	CreateProducts(product []models.Product) error
	GetProducts(types.GetCategoriesParams) []models.Product
}

type IProductService interface {
	CreateProducts(products []types.ProductRequest) error
	GetProducts(types.GetCategoriesParams) ([]types.ProductRequest, error)
}
