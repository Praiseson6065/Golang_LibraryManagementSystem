package payments

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type PaymentRequest struct {
	Amount        float64 `json:"amount"`
	PaymentMethod string  `json:"paymentMethod"`
	UserId        uint    `json:"userId"`
	Email         string  `json:"email"`
	Books         []Book  `json:"books"`
}

func makePayment() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var paymentRequest PaymentRequest
		if err := ctx.ShouldBindBodyWithJSON(&paymentRequest); err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		sessionId, err := createCheckoutSession(&paymentRequest)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		payment := Payment{
			Amount:        paymentRequest.Amount,
			PaymentMethod: paymentRequest.PaymentMethod,
			UserId:        paymentRequest.UserId,
			Email:         paymentRequest.Email,
			Books:         paymentRequest.Books,
			Status:        "Pending",
		}
		err = createPayment(ctx, &payment)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"sessionId": sessionId})

	}
}
