package auth

import "github.com/gin-gonic/gin"

func firstOrCreateUser(ctx *gin.Context, user *User) (string, error) {
	tx := db.WithContext(ctx).Create(user)
	if tx.Error != nil {
		return "", tx.Error
	}
	return string(user.ID), tx.Error
}

func getPasswordAndRole(ctx *gin.Context, email string) (string, uint, uint, error) {
	var user User
	tx := db.WithContext(ctx).Where("email = ?", email).First(&user)
	return user.Password, user.ID, uint(user.RoleId), tx.Error
}
