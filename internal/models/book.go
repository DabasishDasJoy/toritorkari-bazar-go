package models

type Book struct {
	ID          uint `gorm:"primaryKey;autoIncrement"`
	BookName    string
	Author      string
	Publication string
}
