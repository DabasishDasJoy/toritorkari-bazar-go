package domain

import (
	"toritorkari-bazar/internal/models"
	"toritorkari-bazar/types"
)

type IProductRepo interface {
	CreateProducts(product []models.Product) error
}

type IProductService interface {
	CreateProducts(products []types.ProductRequest) error
}
