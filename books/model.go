package books

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	ID          uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	Title       string    `json:"title" binding:"required"`
	Author      string    `json:"author" binding:"required"`
	Description string    `json:"description" binding:"required"`
	Quantity    uint      `json:"quantity" binding:"required"`
	Price       uint      `json:"price" binding:"required"`
	ImageUrl    string    `json:"imageUrl" binding:"required"`
}
