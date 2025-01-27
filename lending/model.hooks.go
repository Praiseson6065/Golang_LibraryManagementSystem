package lending

import (
	"strings"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (obj *LendBook) BeforeCreate(tx *gorm.DB) (err error) {
	obj.ID = "L" + strings.Replace(uuid.New().String(), "-", "", -1)
	return
}
