package service

import (
	"errors"
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
			SubCategoryID: product.SubCategoryID,
			CategoryID:    product.CategoryID,
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

func (service ProductService) GetProducts(getCategoriesParams types.GetCategoriesParams) (types.ProductResponse, error) {
	productsResponse, total, _ := service.repo.GetProducts(getCategoriesParams)
	var productResponse = types.ProductResponse{}

	if len(productsResponse) == 0 {
		return productResponse, errors.New("no products found")
	}

	var products []types.ProductRequest

	for _, product := range productsResponse {
		products = append(products, types.ProductRequest{
			ID:            product.ID,
			Name:          product.Name,
			Description:   product.Description,
			Price:         product.Price,
			Quantity:      product.Quantity,
			SubCategoryID: product.SubCategoryID,
			CategoryID:    product.CategoryID,
			Icon:          product.Icon,
			Discount:      product.Discount,
			Status:        product.Status,
		})
	}

	return types.ProductResponse{
		Products: products,
		Count:    total,
	}, nil
}

func (service ProductService) GetProduct(id uint) (types.ProductRequest, error) {
	productResponse, err := service.repo.GetProduct(id)

	if err != nil {
		return types.ProductRequest{}, err
	}

	return types.ProductRequest{
		ID:            productResponse.ID,
		Name:          productResponse.Name,
		Description:   productResponse.Description,
		Price:         productResponse.Price,
		Quantity:      productResponse.Quantity,
		SubCategoryID: productResponse.SubCategoryID,
		CategoryID:    productResponse.CategoryID,
		Icon:          productResponse.Icon,
		Discount:      productResponse.Discount,
		Status:        productResponse.Status,
	}, nil
}
