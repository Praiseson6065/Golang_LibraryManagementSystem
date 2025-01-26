package books

import (
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	ID          string `gorm:"primary_key" gorm:"<-create"`
	Title       string `json:"title" binding:"required"`
	Author      string `json:"author" binding:"required"`
	Description string `json:"description" binding:"required"`
	Quantity    uint   `json:"quantity" binding:"required"`
	Price       uint   `json:"price" binding:"required"`
	ImageUrl    string `json:"imageUrl" binding:"required"`
}
