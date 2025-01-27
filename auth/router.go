package auth

import (
	"LibManMicroServ/events"

	"github.com/gin-gonic/gin"
)

func Router(eventBus *events.EventBus, r *gin.Engine) {
	auth := r.Group("/auth")
	{
		auth.POST("/signup", userSignup(eventBus))
		auth.POST("/login", userLogin())
	}
}
