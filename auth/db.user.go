package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func firstOrCreateUser(ctx *gin.Context, user *User) (uuid.UUID, error) {
	tx := db.WithContext(ctx).Create(user)
	if tx.Error != nil {
		return uuid.UUID{}, tx.Error
	}
	return user.ID, tx.Error
}

func getPasswordAndRole(ctx *gin.Context, email string) (string, uuid.UUID, uint, error) {
	var user User
	tx := db.WithContext(ctx).Where("email = ?", email).First(&user)
	return user.Password, user.ID, uint(user.RoleId), tx.Error
}
