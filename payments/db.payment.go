package payments

import "github.com/gin-gonic/gin"

func createPayment(ctx *gin.Context, payment *Payment) error {
	tx := db.WithContext(ctx).Create(&payment)

	if tx.Error != nil {

		return tx.Error

	}

	return nil

}
