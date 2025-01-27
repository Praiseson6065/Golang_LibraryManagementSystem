package lending

import (
	"github.com/gin-gonic/gin"
)

func LendABook(ctx *gin.Context, lendBook LendBook) (string, error) {
	tx := db.WithContext(ctx).Create(&lendBook)
	if tx.Error != nil {
		return "", tx.Error
	}
	return lendBook.ID, nil
}

func LendBooks(ctx *gin.Context, lendBooks []LendBook) ([]string, error) {
	tx := db.WithContext(ctx).Create(&lendBooks)
	if tx.Error != nil {
		return nil, tx.Error
	}
	var ids []string
	for _, lendBook := range lendBooks {
		ids = append(ids, lendBook.ID)
	}
	return ids, nil
}

func ReturnABook(ctx *gin.Context, lendBook LendBook) error {
	tx := db.WithContext(ctx).Updates(&lendBook)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func ReturnBooks(ctx *gin.Context, lendBook []LendBook) error {
	tx := db.WithContext(ctx).Updates(&lendBook)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func GetLendBooksByUserId(ctx *gin.Context, UserId string) ([]LendBook, error) {
	var lendBook []LendBook
	tx := db.WithContext(ctx).Where("user_id = ?", UserId).Find(&lendBook)
	if tx.Error != nil {
		return lendBook, tx.Error
	}
	return lendBook, nil
}
