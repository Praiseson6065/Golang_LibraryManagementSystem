package cart

import (
	"LibManMicroServ/events"

	"github.com/gin-gonic/gin"
)

func Router(eventBus *events.EventBus, r *gin.Engine) {

	userCartEventsQueue := eventBus.Subscribe(string(events.EventUserSignedUp))

	go func() {
		for event := range userCartEventsQueue {
			payload := event.Payload.(events.EventUserSignedUpPayload)
			c := event.Context
			CreateUserCart(c, payload.UserId)
		}
	}()
		
}
