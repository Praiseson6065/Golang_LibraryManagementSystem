package cart

import (
	"LibManMicroServ/events"
	"LibManMicroServ/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CheckOutCart(eventBus *events.EventBus) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userId := middleware.GetUserID(ctx)
		getCart, err := GetCart(ctx, userId)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		var validationLendItems []events.BookForPayload
		var validationCartItems []events.BookForPayload

		for _, item := range getCart.LendItems {
			validationLendItems = append(validationLendItems, events.BookForPayload{item.BookID, 1, false})
		}
		for _, item := range getCart.PurchaseItems {
			validationCartItems = append(validationCartItems, events.BookForPayload{item.BookID, item.Quantity, false})
		}
		eventBus.Publish(events.Event{
			Type:    events.EventBooksValidation,
			Context: ctx,
			Payload: events.EventCartCheckedOutPayload{
			},
		})
		

	}
}
