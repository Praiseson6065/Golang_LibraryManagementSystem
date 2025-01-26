package books

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func bookAdd() gin.HandlerFunc {

	return func(ctx *gin.Context) {

		var bookAddRequest Book
		if err := ctx.ShouldBindBodyWithJSON(&bookAddRequest); err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		bookId, err := addBook(ctx, bookAddRequest)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusCreated, gin.H{"bookId": bookId})

	}

}

func bookUpdate() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		bookId := ctx.Param("id")
		var bookUpdateRequest Book
		if err := ctx.ShouldBindBodyWithJSON(&bookUpdateRequest); err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		_, err := updateBook(ctx, bookUpdateRequest)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

			return
		}

		ctx.JSON(http.StatusOK, gin.H{"bookId": bookId})
	}

}

func bookDelete() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		bookId := ctx.Param("id")

		err := deleteBook(ctx, bookId)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"bookId": bookId})
	}
}
