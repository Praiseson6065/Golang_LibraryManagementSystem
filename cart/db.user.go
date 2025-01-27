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
