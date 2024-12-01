package controllers

import (
	"log"
	"net/http"
	"toritorkari-bazar/pkg/domain"
	"toritorkari-bazar/pkg/models"
	"toritorkari-bazar/pkg/types"

	"github.com/labstack/echo/v4"
)

var CategoryService domain.ICategoryService

func SetCategoryService(categoryService domain.ICategoryService) {
	CategoryService = categoryService
}

func GetCategories(e echo.Context) error {
	categories, err := CategoryService.GetCategories()

	if err != nil {
		return e.JSON(http.StatusInternalServerError, err.Error())
	}

	return e.JSON(http.StatusFound, categories)
}

func CreateCategories(e echo.Context) error {
	reqCategories := &[]types.CategoryRequest{}

	if err := e.Bind(reqCategories); err != nil {
		log.Printf("Bind error: %v", err)
		return e.JSON(http.StatusBadRequest, map[string]string{
			"message": "Invalid Data Biding",
		})
	}

	validationErrors := map[int]string{}

	for i, category := range *reqCategories {
		if err := category.ValidateCategory(); err != nil {
			validationErrors[i] = err.Error()
		}
	}

	if len(validationErrors) > 0 {
		return e.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Invalid data",
			"errors":  validationErrors,
		})
	}

	categories := make([]*models.Category, 0, len(*reqCategories))
	for _, category := range *reqCategories {
		categories = append(categories, &models.Category{
			Name: category.Name,
			Icon: category.Icon,
		})
	}

	if err := CategoryService.CreateCategories(categories); err != nil {
		return e.JSON(http.StatusInternalServerError, err.Error())
	}

	return e.JSON(http.StatusOK, map[string]string{
		"message": "categories created successfully",
	})
}
