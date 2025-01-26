package auth

import (
	"LibManMicroServ/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func userLogin() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var loginRequest LoginRequest
		if err := ctx.ShouldBindBodyWithJSON(&loginRequest); err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		hashedPwd, userId, roleId, err := getPasswordAndRole(ctx, loginRequest.Email)

		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		if !comparePasswords(hashedPwd, loginRequest.Password) {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Invalid password",
			})
			return
		}
		token, err := middleware.GenerateToken(string(userId), roleId)

		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{
			"token": token,
		})
	}
}
