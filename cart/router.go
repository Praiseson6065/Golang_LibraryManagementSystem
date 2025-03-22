package cart

import (
	"LibManMicroServ/events"
	_ "LibManMicroServ/docs/usercart"
	"github.com/gin-gonic/gin"
)

func Router(eventBus *events.EventBus, r *gin.Engine) {

	userCartEventsQueue, responses := eventBus.Subscribe(string(events.EventUserSignedUp))

	go func() {
		for event := range userCartEventsQueue {
			payload := event.Payload.(events.EventUserSignedUpPayload)
			c := event.Context

			responses <- EventCreateUserCart(c, payload.UserId)
		}
	}()
	cart := r.Group("/cart")
	{
		cart.POST("/checkout", CheckOutCart(eventBus))
		cart.POST("/addItemToCart", AddItemToCart(eventBus))
	}

}
