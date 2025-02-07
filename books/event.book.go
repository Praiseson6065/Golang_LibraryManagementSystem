package books

import (
	"LibManMicroServ/events"

	"github.com/gin-gonic/gin"
)

func validateCheckOutCart(ctx *gin.Context, payload events.EventCartCheckedOutPayload) (events.EventCartCheckedOutPayload, bool, error) {

	flag := true

	for _, item := range payload.LendCartItem {
		checkAvailability, err := IsBookQuantityAvailable(ctx, item.BookId, uint(item.Quantity))
		item.Available = checkAvailability
		flag = flag && checkAvailability
		if err != nil {
			return payload, false, err
		}

	}

	return payload, flag, nil

}
