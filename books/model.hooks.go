package books

import (
	"strings"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (book *Book) BeforeCreate(tx *gorm.DB) (err error) {
	book.ID = "B" + strings.Replace(uuid.New().String(), "-", "", -1)
	return
}
