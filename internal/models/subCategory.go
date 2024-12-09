package models

type SubCategory struct {
	ID         uint   `gorm:"primaryKey;autoIncrement"`
	Name       string `gorm:"unique"`
	CategoryId uint   `gorm:"not null"`
}
