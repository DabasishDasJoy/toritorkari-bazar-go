package controllers

import (
	"log"
	"net/http"
	"strconv"

	"toritorkari-bazar/internal/domain"
	"toritorkari-bazar/types"

	"github.com/labstack/echo/v4"
)

var CategoryService domain.ICategoryService

func SetCategoryService(categoryService domain.ICategoryService) {
	CategoryService = categoryService
}

func GetCategories(e echo.Context) error {
	tempCategoryId := e.QueryParam("categoryId")
	categoryId, err := strconv.ParseInt(tempCategoryId, 0, 0)

	if err != nil && tempCategoryId != "" {
		return e.JSON(http.StatusBadRequest, "Invalid category ID")
	}

	categories, err := CategoryService.GetCategories(uint(categoryId))

	if err != nil {
		return e.JSON(http.StatusInternalServerError, err.Error())
	}

	return e.JSON(http.StatusFound, categories)
}

func CreateCategories(e echo.Context) error {
	reqCategories := []types.CategoryRequest{}

	if err := e.Bind(&reqCategories); err != nil {
		log.Printf("Bind error: %v", err)
		return e.JSON(http.StatusBadRequest, map[string]string{
			"message": "Invalid Data Biding",
		})
	}

	validationErrors := map[int]string{}

	for i, category := range reqCategories {
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

	if err := CategoryService.CreateCategories(reqCategories); err != nil {
		return e.JSON(http.StatusInternalServerError, err.Error())
	}

	return e.JSON(http.StatusOK, map[string]string{
		"message": "categories created successfully",
	})
}
