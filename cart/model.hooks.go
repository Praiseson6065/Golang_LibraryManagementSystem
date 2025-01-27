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

