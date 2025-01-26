package auth

import "github.com/gin-gonic/gin"

func Router(r *gin.Engine) {
	auth := r.Group("/auth")
	{
		auth.POST("/signup", userSignup())
		auth.POST("/login", userLogin())
	}
}
