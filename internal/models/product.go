package models

type Product struct {
	ID            uint   `gorm:"primaryKey;autoIncrement"`
	Name          string `gorm:"unique"`
	Description   string
	CategoryId    uint    `gorm:"not null"`
	SubCategoryId uint    `gorm:"not null"`
	Icon          string  `gorm:"not null"`
	Price         float64 `gorm:"not null"`
	Quantity      string  `gorm:"not null"`
	Discount      int     `gorm:"not null"`
	Status        string  `gorm:"not null"`
}
