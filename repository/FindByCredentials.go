package repository

import (
	"database/sql"
	"errors"

	"github.com/Praiseson6065/Golang_LibraryManagementSystem/middlewares"
	"github.com/Praiseson6065/Golang_LibraryManagementSystem/models"
	_ "github.com/lib/pq"
)

func FindByCredentials(email, password string) (*models.User, error) {

	// Connect to the database
	db, err := sql.Open("postgres", "postgres://lib:lib@localhost:5432/lib")
	if err != nil {
		return nil, err
	}
	defer db.Close()

	// Build the query
	query := `SELECT id, email, password, name FROM user_data WHERE email = $1 ;`

	// Execute the query

	if err != nil {
		return nil, err
	}
	result, err := db.Query(query, email)
	if err != nil {
		return nil, err
	}

	// Check if the query returned any rows
	if !result.Next() {
		return nil, errors.New("user not found")
	}
	user := models.User{}
	err = result.Scan(&user.ID, &user.Email, &user.Password, &user.Name)
	if err != nil {
		return nil, err
	}

	// Check the password
	matched := middlewares.CheckPasswordHash(password,user.Password)

	// If the password matches, return the user
	if matched {
		return &user, nil
	} else {
		return nil, errors.New("invalid password")
	}
}
