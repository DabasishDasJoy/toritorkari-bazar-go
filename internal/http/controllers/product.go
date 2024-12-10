package controllers

import (
	"log"
	"net/http"
	"strconv"
	"toritorkari-bazar/internal/domain"
	"toritorkari-bazar/types"

	"github.com/labstack/echo/v4"
)

var ProductService domain.IProductService

func ProductServiceInstance(productService domain.IProductService) {
	ProductService = productService
}

func CreateProducts(e echo.Context) error {
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

		if Category, err := CategoryService.GetCategories(product.CategoryId); err != nil {
			log.Printf("Get Categories Error: %v", err)
			validationErrors[i] = "invalid category id " + strconv.Itoa(int(product.CategoryId))
		} else {
			log.Print("valid category with id", strconv.Itoa(int(Category[0].ID)))
		}

		if SubCategory, err := SubCategoryService.GetSubCategory(product.SubCategoryId); err != nil {
			log.Printf("Get Sub Categories Error: %v", err)
			validationErrors[i] = "invalid subcategory id " + strconv.Itoa(int(product.SubCategoryId))
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

	if err := ProductService.CreateProducts(reqProducts); err != nil {
		return e.JSON(http.StatusInternalServerError, err.Error())
	}

	return e.JSON(http.StatusOK, map[string]string{
		"message": "products created successfully",
	})
}

func GetProducts(e echo.Context) error {
	temporaryCategoryId := e.QueryParam("categoryId")
	categoryId, err := strconv.ParseUint(temporaryCategoryId, 0, 0)

	if err != nil {
		e.JSON(http.StatusBadRequest,
			"invalid category id",
		)
	}

	products, err := ProductService.GetProducts(uint(categoryId))

	if err != nil {
		e.JSON(http.StatusInternalServerError, err.Error())
	}

	return e.JSON(http.StatusFound, products)
}
