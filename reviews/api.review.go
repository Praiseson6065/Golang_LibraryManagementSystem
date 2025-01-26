package reviews

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func getReviewsByBookID() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("bookId")
		reviews, err := GetAllReviewsByBookID(ctx, id)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, reviews)
	}
}
