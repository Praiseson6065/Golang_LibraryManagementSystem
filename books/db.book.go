package books

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func getBooks(ctx *gin.Context) ([]Book, error) {
	var books []Book
	tx := db.WithContext(ctx).Find(&books)
	if tx.Error != nil {
		return books, tx.Error
	}
	return books, nil
}

func getBook(ctx *gin.Context, id string) (Book, error) {
	var book Book
	tx := db.WithContext(ctx).First(&book, id)
	if tx.Error != nil {
		return book, tx.Error
	}
	return book, nil
}

func addBook(ctx *gin.Context, book Book) (uuid.UUID, error) {
	tx := db.WithContext(ctx).Create(&book)
	if tx.Error != nil {
		return uuid.UUID{}, tx.Error
	}
	return book.ID, nil
}

func updateBook(ctx *gin.Context, book Book) (uuid.UUID, error) {
	tx := db.WithContext(ctx).Save(&book)
	if tx.Error != nil {
		return uuid.UUID{}, tx.Error
	}
	return book.ID, nil
}

func deleteBook(ctx *gin.Context, id uuid.UUID) error {
	tx := db.WithContext(ctx).Delete(&Book{}, id)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}
