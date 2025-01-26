package payments

import "gorm.io/gorm"

type Book struct {
	gorm.Model
	BookId   uint    `json:"bookId"`
	Title	string  `json:"title"`
	Quantity uint    `json:"quantity"`
	Price    float64 `json:"price"`
	ImageURL string  `json:"imageURL"`
}

type Payment struct {
	gorm.Model
	Amount            float64 `json:"amount"`
	PaymentMethod     string  `json:"paymentMethod"`
	UserId            uint    `json:"userId"`
	Email             string  `json:"email"`
	Books             []Book  `json:"books" gorm:"many2many:payment_books;"`
	CheckoutSessionID string  `json:"checkout_session_id"`
	Status            string  `json:"status"`
}
