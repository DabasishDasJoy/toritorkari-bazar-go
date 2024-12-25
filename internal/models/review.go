package models

import (
	"time"
)

type Review struct {
	ID        uint    `gorm:"primaryKey;autoIncrement"`
	Rating    float64 `gorm:"notnull"`
	Review    string  `gorm:"not null"`
	ProductID uint    `gorm:"foreignKey:Product"`
	UserID    uint    `gorm:"foreignKey:User"`
	CreatedAt time.Time
}

type ReviewResponse struct {
	Review
	Email string
	Name  string
}
