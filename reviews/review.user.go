package reviews

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func getAllReviewsByUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		reviews, err := GetAllReviewsByUserID(ctx)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, reviews)

	}
}

func writeReview() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var review Review
		if err := ctx.ShouldBindJSON(&review); err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		reviewId, err := CreateAReview(ctx, review)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"reviewId": reviewId})
	}
}

func updateReview() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var review Review
		if err := ctx.ShouldBindJSON(&review); err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		reviewId, err := UpdateAReview(ctx, review)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"reviewId": reviewId})
	}
}

func deleteReview() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		reviewId := ctx.Param("id")

		err := DeleteAReview(ctx, reviewId)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"message": "Review Deleted"})
	}
}
