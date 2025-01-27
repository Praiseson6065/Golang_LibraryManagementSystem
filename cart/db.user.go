package cart

import (
	"LibManMicroServ/middleware"

	"github.com/gin-gonic/gin"
)

func CreateUserCart(ctx *gin.Context) error {
	var cart Cart
	cart.UserID = middleware.GetUserID(ctx)
	tx := db.WithContext(ctx).Create(&cart)

	if tx.Error != nil {

		return tx.Error

	}
	return nil
}
