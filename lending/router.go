package lending

import "github.com/gin-gonic/gin"

func UserRouter(r *gin.Engine) {
	r.Group("/lending")
	{
		r.POST("/lend", BookLending())
		r.PUT("/return", BookReturning())

	}
}

func AdminRouter(r *gin.Engine) {
	r.Group("/admin/lending")
	{
		r.POST("/approve", ApproveBooks())
	}
}
