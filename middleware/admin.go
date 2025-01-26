package middleware

import (
	"LibManMicroServ/constants"
	"net/http"

	"github.com/gin-gonic/gin"
)

func EnsureAdmin() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		RoleId := GetRoleID(ctx)

		if RoleId != uint(constants.Librarian) {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "You are not authorized to access this resource"})
			return
		}

		ctx.Next()

	}
}
