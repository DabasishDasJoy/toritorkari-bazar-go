package controllers

import (
	"log"
	"net/http"
	"toritorkari-bazar/pkg/domain"
	"toritorkari-bazar/pkg/models"
	"toritorkari-bazar/pkg/types"

	"github.com/labstack/echo/v4"
)

var SubCategoryService domain.ISubCategoryService

func SetSubCategoryService(subCategoryService domain.ISubCategoryService) {
	SubCategoryService = subCategoryService
}

func CreateSubCategories(e echo.Context) error {
	reqSubCategory := &[]types.SubCategoryRequest{}

	if err := e.Bind(reqSubCategory); err != nil {
		log.Printf("Bind Error: %v", err)

		return e.JSON(http.StatusBadRequest, map[string]string{
			"message": "invalid data binding",
		})
	}

	validationErrors := map[int]string{}

	for i, subCategory := range *reqSubCategory {
		if err := subCategory.ValidateSubCategory(); err != nil {
			validationErrors[i] = err.Error()
		}
	}

	if len(validationErrors) > 0 {
		return e.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Invalid data",
			"errors":  validationErrors,
		})
	}

	subCategories := make([]*models.SubCategory, 0, len(*reqSubCategory))

	for _, subCategory := range *reqSubCategory {
		subCategories = append(subCategories, &models.SubCategory{
			Name:       subCategory.Name,
			CategoryId: subCategory.CategoryId,
		})
	}

	if err := SubCategoryService.CreateSubCategories(subCategories); err != nil {
		return e.JSON(http.StatusInternalServerError, err.Error())
	}

	return e.JSON(http.StatusOK, map[string]string{
		"message": "Sub Categories created successfully",
	})
}
