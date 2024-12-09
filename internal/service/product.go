package service

import (
	"toritorkari-bazar/internal/domain"
	"toritorkari-bazar/internal/models"
	"toritorkari-bazar/types"
)

type ProductService struct {
	repo domain.IProductRepo
}

func ProductServiceInstance(productRepo domain.IProductRepo) domain.IProductService {
	return &ProductService{
		repo: productRepo,
	}
}

func (service ProductService) CreateProducts(reqProducts []types.ProductRequest) error {
	products := make([]models.Product, 0, len(reqProducts))

	for _, product := range reqProducts {
		products = append(products, models.Product{
			Name:          product.Name,
			Description:   product.Description,
			Price:         product.Price,
			Quantity:      product.Quantity,
			SubCategoryId: product.SubCategoryId,
			CategoryId:    product.CategoryId,
			Icon:          product.Icon,
			Discount:      product.Discount,
			Status:        product.Status,
		})
	}

	if err := service.repo.CreateProducts(products); err != nil {
		return err
	}

	return nil
}
