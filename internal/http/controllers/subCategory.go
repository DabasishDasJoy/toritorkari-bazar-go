package controllers

import (
	"log"
	"net/http"
	"strconv"

	"toritorkari-bazar/internal/domain"
	"toritorkari-bazar/types"

	"github.com/labstack/echo/v4"
)

type SubCategoryController struct {
	SubCategoryService domain.ISubCategoryService
	CategoryService    domain.ICategoryService
}

func SetSubCategoryService(subCategoryService domain.ISubCategoryService, categoryService domain.ICategoryService) *SubCategoryController {
	return &SubCategoryController{
		SubCategoryService: subCategoryService,
		CategoryService:    categoryService,
	}
}

func (controller *SubCategoryController) CreateSubCategories(e echo.Context) error {
	reqSubCategory := []types.SubCategoryRequest{}

	if err := e.Bind(&reqSubCategory); err != nil {
		log.Printf("Bind Error: %v", err)

		return e.JSON(http.StatusBadRequest, map[string]string{
			"message": "invalid data binding",
		})
	}

	validationErrors := map[int]string{}

	for i, subCategory := range reqSubCategory {
		if err := subCategory.ValidateSubCategory(); err != nil {
			validationErrors[i] = err.Error()
		}

		if categories, err := controller.CategoryService.GetCategories(subCategory.CategoryId); err != nil {
			log.Printf("Get Categories Error: %v", err)
			validationErrors[i] = "invalid category id " + strconv.Itoa(int(subCategory.CategoryId))
		} else {
			log.Print(categories)
		}
	}

	if len(validationErrors) > 0 {
		return e.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Invalid data",
			"errors":  validationErrors,
		})
	}

	if err := controller.SubCategoryService.CreateSubCategories(reqSubCategory); err != nil {
		return e.JSON(http.StatusInternalServerError, err.Error())
	}

	return e.JSON(http.StatusOK, map[string]string{
		"message": "Sub Categories created successfully",
	})
}
