package books

import (
	"LibManMicroServ/events"

	"github.com/gin-gonic/gin"
)

func validateCheckOutCart(ctx *gin.Context, payload events.EventCartCheckedOutPayload) events.EventCartCheckedOutResponse {

	var response events.EventCartCheckedOutResponse
	for _, item := range payload.LendCartItem {
		checkAvailability, availableQuantity, err := IsBookQuantityAvailable(ctx, item.BookId, uint(item.Quantity))
		item.Available = checkAvailability
		item.Quantity = availableQuantity
		response.Success = response.Success && checkAvailability
		response.LendCartItem = append(response.LendCartItem, item)
		if err != nil {
			response.Success = false
			response.Error = err
			return response
		}

	}
	for _, item := range payload.PurchaseCartItem {
		checkAvailability, availableQuantity, err := IsBookQuantityAvailable(ctx, item.BookId, uint(item.Quantity))
		item.Available = checkAvailability
		item.Quantity = availableQuantity
		response.Success = response.Success && checkAvailability
		response.PurchaseCartItem = append(response.PurchaseCartItem, item)
		if err != nil {
			response.Success = false
			response.Error = err
			return response
		}

	}
	response.Success = true
	return response

}
