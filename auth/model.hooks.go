package auth

import (
	"strings"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (user *User) BeforeCreate(tx *gorm.DB) (err error) {
	user.ID = "U" + strings.Replace(uuid.New().String(), "-", "", -1)
	return
}
