package cart

import (
	"github.com/gin-gonic/gin"
)

func CreateUserCart(ctx *gin.Context, UserId string) error {
	var cart Cart
	cart.UserID = UserId
	tx := db.WithContext(ctx).Create(&cart)

	if tx.Error != nil {

		return tx.Error

	}
	return nil
}

func GetCart(ctx *gin.Context, UserId string) (Cart, error) {
	var cart Cart
	tx := db.WithContext(ctx).Where("user_id = ?", UserId).First(&cart)

	if tx.Error != nil {
		return cart, tx.Error
	}
	return cart, nil
}
