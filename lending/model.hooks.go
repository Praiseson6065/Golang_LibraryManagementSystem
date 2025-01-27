package lending

import (
	"strings"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (obj *LendBook) BeforeCreate(tx *gorm.DB) (err error) {
	obj.ID = "L" + strings.Replace(uuid.New().String(), "-", "", -1)
	return
}

func (obj *LendBook) BeforeUpdate(tx *gorm.DB) (err error) {
	if obj.Status == "Returned" && obj.ReturnedAt == nil {
		now := time.Now()
		obj.ReturnedAt = &now
	} else if obj.Status == "Approved" && obj.IsApproved == false {
		obj.StartDate = time.Now()
		obj.IsApproved = true
	}
	return
}
