package models

type Category struct {
	ID   uint   `gorm:"primaryKey;autoIncrement"`
	Name string `gorm:"unique"`
	Icon string
}

type IntermediateCategoryResponse struct {
	Category
	SubCategoryID   uint
	SubCategoryName string
}
