package repository

import (
	"database/sql"
	"errors"

	"github.com/Praiseson6065/Golang_LibraryManagementSystem/models"
)

func FindByCredentials(email, password string) (*models.User, error) {

	// Connect to the database
	db, err := sql.Open("postgres", "lib:lib@tcp(localhost:5432)/lib")
	if err != nil {
		return nil, err
	}
	defer db.Close()

	// Build the query
	query := `SELECT id, email, password, favph FROM user_data WHERE email = $1 AND password = $2;`

	// Execute the query
	result, err := db.Query(query, email, password)
	if err != nil {
		return nil, err
	}

	// Check if the query returned any rows
	if !result.Next() {
		return nil, errors.New("user not found")
	}

	// Scan the row into a User struct
	user := models.User{}
	err = result.Scan(&user.ID, &user.Email, &user.Password, &user.FavoritePhrase)
	if err != nil {
		return nil, err
	}

	// Return the user
	return &user, nil
}
