package books

import "github.com/gin-gonic/gin"

func Router(r *gin.Engine) {
	books := r.Group("/books")
	{
		books.GET("/", getAllBooks())
		books.GET("/:id", getOneBook())
	}
}

func AdminRouter(r *gin.Engine) {
	books := r.Group("/")
	{
		books.POST("/", bookAdd())
		books.PUT("/", bookUpdate())
		books.DELETE("/:id", bookDelete())
	}
}
