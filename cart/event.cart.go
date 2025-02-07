package cart

import (
	"LibManMicroServ/events"

	"github.com/gin-gonic/gin"
)

func EventCreateUserCart(c *gin.Context, UserId string) events.EventUserSignedUpResponse {

	err := CreateUserCart(c, UserId)
	if err != nil {
		return events.EventUserSignedUpResponse{false, err}
	}
	return events.EventUserSignedUpResponse{true, nil}

}
