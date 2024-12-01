package types

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type BookRequest struct {
	ID          uint   `json:"bookId"`
	BookName    string `json:"bookname"`
	Author      string `json:"author"`
	Publication string `json:"publication,omitempty"`
}

type CategoryRequest struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
	Icon string `json:"icon"`
}

type SubCategoryRequest struct {
	ID         uint   `json:"id"`
	Name       string `json:"name"`
	CategoryId uint   `json:"categoryId"`
}

func (book BookRequest) Validate() error {
	return validation.ValidateStruct(&book,
		validation.Field(&book.BookName,
			validation.Required.Error("Book name cannot be empty"),
			validation.Length(1, 50)),
		validation.Field(&book.Author,
			validation.Required.Error("Author name cannot be empty"),
			validation.Length(5, 50)),
	)
}

func (category CategoryRequest) ValidateCategory() error {
	return validation.ValidateStruct(&category,
		validation.Field(&category.Name,
			validation.Required.Error("Category name cannot be empty"),
			validation.Length(1, 50)),
		validation.Field(&category.Icon,
			validation.Required.Error("Icon cannot be empty")),
	)
}

func (subCategory SubCategoryRequest) ValidateSubCategory() error {
	return validation.ValidateStruct(&subCategory,
		validation.Field(&subCategory.Name,
			validation.Required.Error("Category name cannot be empty"),
			validation.Length(1, 50)),
		validation.Field(&subCategory.CategoryId,
			validation.Required.Error("CategoryId cannot be empty")),
	)
}
