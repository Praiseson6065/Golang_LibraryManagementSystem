package cart

import (
	"errors"
	"strings"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (obj *Cart) BeforeCreate(tx *gorm.DB) (err error) {
	obj.ID = "C" + strings.Replace(uuid.New().String(), "-", "", -1)
	if obj.UserID == "" {
		return errors.New("user ID is required")
	}
	return
}
func (lc *LendCartItem) BeforeCreate(tx *gorm.DB) (err error) {
	lc.ID = "LC" + strings.Replace(uuid.New().String(), "-", "", -1)
	return
}

func (pc *PurchaseCartItem) BeforeCreate(tx *gorm.DB) (err error) {
	pc.ID = "PC" + strings.Replace(uuid.New().String(), "-", "", -1)
	return
}

