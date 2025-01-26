package reviews

import (
	"strings"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (review *Review) BeforeCreate(tx *gorm.DB) (err error) {
	review.ID = "R" + strings.Replace(uuid.New().String(), "-", "", -1)
	return
}
