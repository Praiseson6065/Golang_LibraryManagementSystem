package reviews

import (
	"github.com/gin-gonic/gin"
)

func UserRouter(r *gin.Engine) {

	review := r.Group("/review")
	{
		review.POST("/write", writeReview())
		review.PUT("/update", updateReview())
		review.DELETE("/delete/:id", deleteReview())
		review.GET("/user", getAllReviewsByUser())
	}

}

func Router(r *gin.Engine) {
	review := r.Group("/review")
	{
		review.GET("/:bookId", getReviewsByBookID())
	}
}
