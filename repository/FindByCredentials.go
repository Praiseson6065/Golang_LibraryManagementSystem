package repository

import (
	"errors"

	"github.com/Praiseson6065/Golang_LibraryManagementSystem/database"
	"github.com/Praiseson6065/Golang_LibraryManagementSystem/middlewares"
	"github.com/Praiseson6065/Golang_LibraryManagementSystem/models"
	_ "github.com/lib/pq"
)

func FindByCredentials(email, password string) (*models.User, error) {
	db, err := database.DbGormConnect()
	if err != nil {
		return nil, err
	}
	user := models.User{}
	db.Where("email = ?", email).First(&user)

	if user.ID == 0 {
		return nil, errors.New("user not found")
	}

	if !middlewares.CheckPasswordHash(password, user.Password) {
		return nil, errors.New("invalid password")
	}

	return &user, nil
}
