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

func (service ProductService) GetProducts(getCategoriesParams types.GetCategoriesParams) (types.ProductListResponse, error) {
	productsResponse, total, _ := service.repo.GetProducts(getCategoriesParams)
	var productResponse = types.ProductListResponse{}

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

	return types.ProductListResponse{
		Products: products,
		Count:    total,
	}, nil
}

func (service ProductService) GetProduct(id uint) (types.ProductResponse, error) {
	productResponse, err := service.repo.GetProduct(id)

	if err != nil {
		return types.ProductResponse{}, err
	}

	return types.ProductResponse{
		ID:          productResponse.ID,
		Name:        productResponse.Name,
		Description: productResponse.Description,
		Price:       productResponse.Price,
		Quantity:    productResponse.Quantity,
		Icon:        productResponse.Icon,
		Discount:    productResponse.Discount,
		Status:      productResponse.Status,
		Category: types.CategoryRequest{
			ID:   productResponse.CategoryID,
			Name: productResponse.CategoryName,
			Icon: productResponse.CategoryIcon,
		},
		SubCategory: types.SubCategoryRequest{
			ID:   productResponse.SubCategoryID,
			Name: productResponse.SubCategoryName,
		},
	}, nil
}
