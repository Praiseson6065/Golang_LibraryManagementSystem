package payments

import "github.com/gin-gonic/gin"

func Router(r *gin.Engine) {
	payments := r.Group("/api/payments")
	{

		payments.POST("/make", makePayment())

	}
}
