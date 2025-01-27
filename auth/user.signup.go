package auth

import (
	"LibManMicroServ/events"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserSignupRequest struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func userSignup(eventsBus *events.EventBus) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var userSignupRequest UserSignupRequest

		if err := ctx.ShouldBindBodyWithJSON(&userSignupRequest); err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		hashedPwd := hashAndSalt(userSignupRequest.Password)

		id, err := firstOrCreateUser(ctx, &User{
			Name:     userSignupRequest.Name,
			Email:    userSignupRequest.Email,
			Password: hashedPwd,
		})

		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		eventsBus.Publish(events.Event{
			Type:    events.EventUserSignedUp,
			Context: ctx,
			Payload: events.EventUserSignedUpPayload{
				UserId: id,
			},
		})

		ctx.JSON(http.StatusOK, gin.H{"status": "Successfully signed up", "userId": id})

	}
}
