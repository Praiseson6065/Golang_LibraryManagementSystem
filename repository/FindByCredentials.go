package repository

import (
	"database/sql"
	"errors"

	"github.com/Praiseson6065/Golang_LibraryManagementSystem/models"
)

func FindByCredentials(db *sql.DB, email, password string) (*models.User, error) {

	query := `SELECT id, email, password, favph FROM user_data WHERE email = ? AND password = ?;`

	result, err := db.Query(query, email, password)
	if err != nil {
		return nil, err
	}

	if result.Next() {
		var user models.User
		err = result.Scan(&user.ID, &user.Email, &user.Password, &user.FavoritePhrase)
		if err != nil {
			return nil, err
		}
		return &user, nil
	} else {
		return nil, errors.New("user not found")
	}
}
