package cart

import (
	"time"

	"gorm.io/gorm"
)

type Cart struct {
	gorm.Model
	ID            string             `json:"id" gorm:"primaryKey"`
	UserID        string             `json:"userId" binding:"required" gorm:"not null"`
	LendItems     []LendCartItem     `json:"lendItems" gorm:"foreignKey:CartID"`
	PurchaseItems []PurchaseCartItem `json:"purchaseItems" gorm:"foreignKey:CartID"`
	CreatedAt     time.Time          `json:"createdAt"`
	UpdatedAt     time.Time          `json:"updatedAt"`
	Status        string             `json:"status" gorm:"type:varchar(20);default:'active'"`
}
type LendCartItem struct {
	gorm.Model
	CartID   string    `json:"cartId" gorm:"not null"`
	BookID   string    `json:"bookId" binding:"required" gorm:"not null"`
	Title    string    `json:"title"`
	ISBN     string    `json:"isbn"`
	LendDate time.Time `json:"lendDate"`
	DueDate  time.Time `json:"dueDate"`
}

type PurchaseCartItem struct {
	gorm.Model
	CartID     string  `json:"cartId" gorm:"not null"`
	BookID     string  `json:"bookId" binding:"required" gorm:"not null"`
	Title      string  `json:"title"`
	ISBN       string  `json:"isbn"`
	Quantity   int     `json:"quantity" gorm:"default:1"`
	UnitPrice  float64 `json:"unitPrice"`
	TotalPrice float64 `json:"totalPrice"`
}
