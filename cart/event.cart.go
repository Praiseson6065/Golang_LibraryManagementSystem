package cart

import (
	"LibManMicroServ/events"

	"github.com/gin-gonic/gin"
)

func EventCreateUserCart(c *gin.Context, UserId string) events.EventUserCartCreationResponse {
	var resp events.EventUserCartCreationResponse
	err := CreateUserCart(c, UserId)
	if err != nil {
		resp.Error = err
		return resp
	}
	return resp

}

func EventValidateBooks(ctx *gin.Context, eB *events.EventBus, lc []LendCartItem, pc []PurchaseCartItem) ([]interface{}, error) {
	var validationLendItems []events.BookForPayload
	var validationCartItems []events.BookForPayload

	for _, item := range lc {
		validationLendItems = append(validationLendItems, events.BookForPayload{BookId: item.BookID, Quantity: 1, Available: false})
	}
	for _, item := range pc {
		validationCartItems = append(validationCartItems, events.BookForPayload{BookId: item.BookID, Quantity: uint(item.Quantity), Available: false})
	}

	responses := eB.Publish(events.Event{
		Type:    events.EventBooksValidation,
		Context: ctx,
		Payload: events.EventCartCheckedOutPayload{},
	})
	return responses, nil

}
