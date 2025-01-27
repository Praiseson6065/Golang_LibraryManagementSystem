package lending

import "github.com/gin-gonic/gin"

type BookLendingRequest struct {
	BookID string `json:"bookId" binding:"required"`
	UserID string `json:"userId" binding:"required"`
	
}

func BookLending() gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}

}
