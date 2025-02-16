package models

type Product struct {
	ID            uint   `gorm:"primaryKey;autoIncrement"`
	Name          string `gorm:"unique"`
	Description   string
	CategoryID    uint    `gorm:"not null"`
	SubCategoryID uint    `gorm:"not null"`
	Icon          string  `gorm:"not null"`
	Price         float64 `gorm:"not null"`
	Quantity      string  `gorm:"not null"`
	Discount      int     `gorm:"not null"`
	Status        string  `gorm:"not null"`
}

type ProductResponse struct {
	ID              uint   `gorm:"primaryKey;autoIncrement"`
	Name            string `gorm:"unique"`
	Description     string
	Icon            string  `gorm:"not null"`
	Price           float64 `gorm:"not null"`
	Quantity        string  `gorm:"not null"`
	Discount        int     `gorm:"not null"`
	Status          string  `gorm:"not null"`
	CategoryName    string
	CategoryIcon    string
	SubCategoryName string
	CategoryID      uint `gorm:"not null"`
	SubCategoryID   uint `gorm:"not null"`
}

type ProductListResponse struct {
	Product
	TotalCount int `json:"totalCount"`
}
