package lending

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ApproveBooks() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var approveBooks []ApproveBooksRequest
		if err := ctx.ShouldBindJSON(&approveBooks); err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		err := ApproveBooksByAdmin(ctx, approveBooks)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"message": "Books Approved successfully"})

	}
}
