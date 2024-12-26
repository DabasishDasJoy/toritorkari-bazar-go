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

type CategoryResponse struct {
	ID            uint                 `json:"id"`
	Name          string               `json:"name"`
	Icon          string               `json:"icon"`
	SubCategories []SubCategoryRequest `json:"subcategories"`
}

type SubCategoryRequest struct {
	ID         uint   `json:"id"`
	Name       string `json:"name"`
	CategoryId uint   `json:"categoryId"`
}

type ProductRequest struct {
	ID            uint    `json:"id"`
	Name          string  `json:"name"`
	Description   string  `json:"description,omitempty"`
	CategoryId    uint    `json:"categoryId"`
	SubCategoryId uint    `json:"subCategoryId"`
	Icon          string  `json:"icon"`
	Price         float64 `json:"price"`
	Quantity      string  `json:"quantity"`
	Discount      int     `json:"discount"`
	Status        string  `json:"status"`
}

type GetCategoriesParams struct {
	CategoryID           uint
	SubCategoryID        uint
	TemporarySearchQuery string
	TemporarySortQuery   string
	Page                 uint
	Size                 uint
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

func (product ProductRequest) ValidateProduct() error {
	return validation.ValidateStruct(&product,
		validation.Field(&product.Name,
			validation.Required.Error("Product name cannot be empty"),
			validation.Length(1, 50)),
		validation.Field(&product.Description,
			validation.Required.Error("Product description cannot be empty")),
		validation.Field(&product.CategoryId,
			validation.Required.Error("CategoryId cannot be empty")),
		validation.Field(&product.SubCategoryId,
			validation.Required.Error("SubCategoryId cannot be empty")),
		validation.Field(&product.Icon,
			validation.Required.Error("Icon cannot be empty")),
		validation.Field(&product.Price,
			validation.Required.Error("Price cannot be empty"),
			validation.Min(0.01).Error("Price must be greater than or equal to 0.01")),
		validation.Field(&product.Quantity,
			validation.Required.Error("Quantity cannot be empty")),
		validation.Field(&product.Discount,
			validation.Min(0).Error("Discount cannot be empty or negative")),
		validation.Field(&product.Status,
			validation.Required.Error("Status cannot be empty"),
			validation.In("in-stock", "stock-out").Error("Status must be either 'in-stock' or 'stock-out'")))
}
