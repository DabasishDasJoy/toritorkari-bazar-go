package models

type User struct {
	ID       uint `gorm:"primaryKey;autoIncrement"`
	Name     string
	Email    string `gorm:"unique"`
	Password string
	Role     string
}
