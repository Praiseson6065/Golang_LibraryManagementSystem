package lending

import (
	"time"

	"gorm.io/gorm"
)

type LendBook struct {
	gorm.Model
	ID         string     `json:"id" gorm:"primaryKey"`
	BookID     string     `json:"bookId" binding:"required" gorm:"not null"`
	UserID     string     `json:"userId" binding:"required" gorm:"not null"`
	AdminID    string     `json:"adminId,omitempty"`
	StartDate  time.Time  `json:"startDate"  gorm:"not null"`
	DueDate    time.Time  `json:"dueDate" gorm:"not null"`
	ReturnedAt *time.Time `json:"returnedAt,omitempty"`
	IsApproved bool       `json:"isApproved" gorm:"default:false"`
	Status     string     `json:"status" gorm:"type:enum('Pending','Approved','Rejected','Returned');default:'Pending'"`
}

type ApproveBooksRequest struct {
	LendBookIDs []string `json:"lendBookIds" binding:"required"`
}
