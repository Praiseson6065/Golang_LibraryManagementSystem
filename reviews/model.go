package reviews

import (
	"gorm.io/gorm"
)

type Review struct {
	gorm.Model
	ID                string `json:"id"`
	BookID            string `json:"bookId"`
	UserID            string `json:"userId"`
	ReviewDescription string `json:"reviewDescription"`
	Rating            int    `json:"rating"`
}
