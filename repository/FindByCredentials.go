package repository

import (
    "errors"
    "github.com/Praiseson6065/Golang_LibraryManagementSystem/models"
)


func FindByCredentials(email, password string) (*models.User, error) {
    
    users := []models.User{
        models.User{
            ID:             1,
            Email:          "test@mail.com",
            Password:       "test12345",
            FavoritePhrase: "Hello, World!",
        },
        models.User{
            ID:             2,
            Email:          "johndoe@mail.com",
            Password:       "password123",
            FavoritePhrase: "I love reading books!",
        },
        models.User{
            ID:             3,
            Email:          "janedoe@mail.com",
            Password:       "password456",
            FavoritePhrase: "I love writing stories!",
        },
    }

    for _, user := range users {
        if user.Email == email && user.Password == password {
            return &user, nil
        }
    }

    return nil, errors.New("user not found")
}
