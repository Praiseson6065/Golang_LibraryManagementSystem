package auth

import (
	"LibManMicroServ/constants"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID       uint           `json:"id" gorm:"primaryKey" gorm:"<-:create"`
	Name     string         `json:"name"`
	RoleId   constants.Role `json:"roleId" gorm:"default:1"`
	Email    string         `json:"email" gorm:"uniqueIndex"`
	Password string         `json:"password"`
}
