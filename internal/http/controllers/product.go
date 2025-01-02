package controllers

import (
	"log"
	"net/http"
	"strconv"
	"toritorkari-bazar/internal/domain"
	"toritorkari-bazar/types"

	"github.com/labstack/echo/v4"
)

type ProductController struct {
	ProductService     domain.IProductService
	SubCategoryService domain.ISubCategoryService
}

func ProductServiceInstance(productService domain.IProductService, subCategoryService domain.ISubCategoryService) *ProductController {
	return &ProductController{
		ProductService:     productService,
		SubCategoryService: subCategoryService,
	}
}

func (controller *ProductController) CreateProducts(e echo.Context) error {
	reqProducts := []types.ProductRequest{}

	if err := e.Bind(&reqProducts); err != nil {
		log.Printf("Bind Error: %v", err)

		return e.JSON(http.StatusBadRequest, map[string]string{
			"message": "invalid data binding",
		})
	}

	// validations
	validationErrors := map[int]string{}

	for i, product := range reqProducts {
		if err := product.ValidateProduct(); err != nil {
			validationErrors[i] = err.Error()
		}

		if Category, err := CategoryService.GetCategories(product.CategoryID); err != nil {
			log.Printf("Get Categories Error: %v", err)
			validationErrors[i] = "invalid category id " + strconv.Itoa(int(product.CategoryID))
		} else {
			log.Print("valid category with id", strconv.Itoa(int(Category[0].ID)))
		}

		if SubCategory, err := controller.SubCategoryService.GetSubCategory(product.SubCategoryID); err != nil {
			log.Printf("Get Sub Categories Error: %v", err)
			validationErrors[i] = "invalid subcategory id " + strconv.Itoa(int(product.SubCategoryID))
		} else {
			log.Print("valid sub category id " + strconv.Itoa(int(SubCategory.ID)))
		}
	}

	if len(validationErrors) > 0 {
		return e.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid data",
			"errors":  validationErrors,
		})
	}

	if err := controller.ProductService.CreateProducts(reqProducts); err != nil {
		return e.JSON(http.StatusInternalServerError, err.Error())
	}

	return e.JSON(http.StatusOK, map[string]string{
		"message": "products created successfully",
	})
}

func (controller *ProductController) GetProducts(e echo.Context) error {
	temporaryCategoryId := e.QueryParam("categoryId")
	temporarySubCategoryId := e.QueryParam("subCategoryId")
	temporarySearchQuery := e.QueryParam("search")
	temporarySortQuery := e.QueryParam("sort")
	temporaryPage := e.QueryParam("page")
	temporarySize := e.QueryParam("size")

	categoryId, err := strconv.ParseUint(temporaryCategoryId, 0, 0)
	if err != nil && temporaryCategoryId != "" {
		return e.JSON(http.StatusBadRequest,
			"invalid category id",
		)
	}

	subCategoryId, err := strconv.ParseUint(temporarySubCategoryId, 0, 0)
	if err != nil && temporarySubCategoryId != "" {
		return e.JSON(http.StatusBadRequest,
			"invalid sub category id",
		)
	}

	page, err := strconv.ParseUint(temporaryPage, 10, 32)
	if err != nil || temporaryPage == "" {
		page = 1
	}

	size, err := strconv.ParseUint(temporarySize, 10, 32)
	if err != nil || temporarySize == "" {
		size = 10
	}

	getProductsParam := types.GetCategoriesParams{
		CategoryID:           uint(categoryId),
		SubCategoryID:        uint(subCategoryId),
		TemporarySearchQuery: temporarySearchQuery,
		TemporarySortQuery:   temporarySortQuery,
		Page:                 uint(page) - 1,
		Size:                 uint(size),
	}

	products, err := controller.ProductService.GetProducts(getProductsParam)

	if err != nil {
		return e.JSON(http.StatusInternalServerError, err.Error())
	}

	return e.JSON(http.StatusFound, products)
}

func (controller *ProductController) GetProduct(e echo.Context) error {
	temporaryProductId := e.Param("productID")
	ProductID, err := strconv.ParseInt(temporaryProductId, 0, 0)

	if err != nil {
		return e.JSON(http.StatusBadRequest, "invalid productID")
	}

	Product, err := controller.ProductService.GetProduct(uint(ProductID))

	if err != nil {
		return e.JSON(http.StatusInternalServerError, err.Error())
	}

	return e.JSON(http.StatusOK, Product)

}
