package lending

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func BookLending() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req []LendBook
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		ids, err := LendBooks(ctx, req)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"ids": ids})

	}

}

func BookReturning() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req []LendBook
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		err := ReturnBooks(ctx, req)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"message": "Books Returned successfully"})

	}

}
