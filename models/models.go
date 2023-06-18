package models

import (
	"fmt"
	"strings"

	"github.com/Praiseson6065/Golang_LibraryManagementSystem/database"
	"gorm.io/gorm"
)

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token string `json:"token"`
}
type User struct {
	gorm.Model
	ID          int    `json:"Id" gorm:"auto_increment:true;primary_key;unique"`
	UserId      string `json:"UserId" gorm:"unique"`
	Name        string `json:"Name"`
	Email       string `json:"Email" gorm:"unique"`
	Password    string `json:"Password"`
	Usertype    string `json:"Usertype"`
	CartBooks   []Book `json:"CartBooks" gorm:"many2many:user_cart_books;"`
	LikedBooks  []Book `json:"LikedBooks" gorm:"many2many:user_liked_books;"`
	IssuedBooks []Book `json:"IssuedBooks" gorm:"many2many:user_issued_books"`
}
type RegisterUser struct {
	Name     string
	Email    string
	Password string
}
type UserData struct {
	ID    int
	Name  string
	Email string
	Exp   int
}
type Logdb struct {
	LogId      int
	UserId     int
	UserType   string
	Operation  string
	InsertTime string
	UserName   string
}

type Book struct {
	gorm.Model
	ID           int    `json:"BookId" gorm:"auto_increment:true;primary_key;unique"`
	BookCode     string `json:"BookCode" gorm:"unique"`
	BookName     string `json:"BookName"`
	ISBN         string `json:"ISBN" gorm:"unique"`
	Pages        int    `json:"Pages"`
	Publisher    string `json:"Publisher"`
	Quantity     int    `json:"Quantity"`
	Author       string `json:"Author"`
	Taglines     string `json:"Taglines"`
	InsertedTime string `json:"InsertedTime"`
	Votes        int    `json:"votes"`
	ImgPath      string `json:"ImgPath"`
}
type SearchBook struct {
	SearchValue  string `json:"SearchValue"`
	SearchColumn string `json:"SearchColumn"`
}

func GetBookByBookCode(BookCode string) (Book, error) {
	db, err := database.DbGormConnect()
	if err != nil {
		return Book{}, err
	}
	book := new(Book)
	db.Where("book_code = ?", BookCode).Find(&book)

	return *book, nil

}
func GetBookById(Id int) (Book, error) {
	db, err := database.DbGormConnect()
	if err != nil {
		return Book{}, err
	}
	book := new(Book)
	db.Where("id = ?", Id).Find(&book)

	return *book, nil

}
func GetBooks() ([]Book, error) {
	db, err := database.DbGormConnect()
	if err != nil {
		fmt.Println(err)
		return []Book{}, err
	}

	books := []Book{}
	db.Find(&books)

	return books, nil

}
func SearchBooks(SearchValue, SearchColumn string) ([]Book, error) {
	var books []Book
	db, err := database.DbGormConnect()
	if err != nil {
		return nil, err
	}
	err = db.Where("lower("+SearchColumn+") like ?", "%"+strings.ToLower(SearchValue)+"%").Find(&books).Error
	if err != nil {
		return nil, err
	}
	return books, nil

}
func GetUser(Id int) (User, error) {
	db, err := database.DbGormConnect()
	if err != nil {
		return User{}, err
	}
	user := new(User)
	db.Where("id=?", Id).Find(&user)
	return *user, nil
}

func GetCartBooksByUserID(userID int) ([]Book, error) {
	db, err := database.DbGormConnect()
	if err != nil {
		return nil, err
	}

	var cartBooks []Book
	if err := db.Joins("JOIN user_cart_books ON user_cart_books.book_id = books.id").
		Where("user_cart_books.user_id = ?", userID).
		Find(&cartBooks).Error; err != nil {
		return nil, err
	}

	return cartBooks, nil
}

