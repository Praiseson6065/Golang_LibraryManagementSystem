package books

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func getAllBooks() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		books, err := getBooks(ctx)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"books": books})
	}

}

func getOneBook() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		bookId := ctx.Param("id")
		book, err := getBook(ctx, bookId)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"book": book})

	}

}
