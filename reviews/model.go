package reviews

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Review struct {
	gorm.Model
	ID                uuid.UUID `json:"id"`
	BookID            uuid.UUID `json:"bookId"`
	UserID            uuid.UUID `json:"userId"`
	ReviewDescription string    `json:"reviewDescription"`
	Rating            int       `json:"rating"`
}
