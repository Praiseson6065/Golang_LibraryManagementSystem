package auth

import (
	"LibManMicroServ/constants"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID       uuid.UUID      `gorm:"primaryKey" gorm:"<-:create"`
	Name     string         `json:"name"`
	RoleId   constants.Role `json:"roleId" gorm:"default:1"`
	Email    string         `json:"email" gorm:"uniqueIndex"`
	Password string         `json:"password"`
}
