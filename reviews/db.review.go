package reviews

import (
	"LibManMicroServ/middleware"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func CreateAReview(ctx *gin.Context, review Review) (uuid.UUID, error) {
	tx := db.WithContext(ctx).Create(&review)
	if tx.Error != nil {

		return uuid.UUID{}, tx.Error

	}
	return review.ID, nil

}

func UpdateAReview(ctx *gin.Context, review Review) (uuid.UUID, error) {
	tx := db.WithContext(ctx).Updates(&review)
	if tx.Error != nil {
		return uuid.UUID{}, tx.Error
	}
	return review.ID, nil
}

func DeleteAReview(ctx *gin.Context, reviewId uuid.UUID) error {
	tx := db.WithContext(ctx).Delete(&Review{}, reviewId)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func GetAllReviewsByUserID(ctx *gin.Context) ([]Review, error) {
	userId := middleware.GetUserID(ctx)
	var reviews []Review
	tx := db.WithContext(ctx).Where("user_id = ?", userId).Find(&reviews)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return reviews, nil
}

func GetAllReviewsByBookID(ctx *gin.Context, bookId string) ([]Review, error) {
	var reviews []Review
	tx := db.WithContext(ctx).Where("book_id = ?", bookId).Find(&reviews)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return reviews, nil
}
