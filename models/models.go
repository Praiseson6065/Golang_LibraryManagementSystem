package models

import (
	"fmt"
	"strconv"
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
type UserRequestedBooks struct {
	UserId        int    `json:"UserId"`
	BookName      string `json:"RequestedBooks"`
	ISBN          string `json:"ISBN"`
	RequestStatus bool   `json:"Status"`
}

type UserData struct {
	ID    int
	Name  string
	Email string
	Exp   int
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
	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}
	sqlDB.Close()
	return *book, nil

}
func GetBookById(Id int) (Book, error) {
	db, err := database.DbGormConnect()
	if err != nil {
		return Book{}, err
	}
	book := new(Book)
	db.Where("id = ?", Id).Find(&book)
	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}
	sqlDB.Close()
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
	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}
	sqlDB.Close()
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
	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}
	sqlDB.Close()
	return books, nil

}
func GetUser(Id int) (User, error) {
	db, err := database.DbGormConnect()
	if err != nil {
		return User{}, err
	}
	user := new(User)
	db.Where("id=?", Id).Find(&user)
	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}
	sqlDB.Close()
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
	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}
	sqlDB.Close()
	return cartBooks, nil
}
func RemovefromCart(UserId, BookId int) (bool, error) {
	db, err := database.DbGormConnect()
	if err != nil {
		return false, err
	}
	query := fmt.Sprintf("Delete from user_cart_books where user_id=%d and book_id=%d ;", UserId, BookId)

	err = db.Exec(query).Error
	if err != nil {
		return false, err
	}
	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}
	sqlDB.Close()
	return true, nil
}
func IssueBooks(Userid int) (bool, error) {
	db, err := database.DbGormConnect()
	if err != nil {
		return false, err
	}
	var cartBooks []Book
	cartBooks, err = GetCartBooksByUserID(Userid)
	if err != nil {
		return false, err
	}
	UserDetails := User{}
	UserDetails, err = GetUser(Userid)
	if err != nil {
		return false, err
	}
	UserDetails.IssuedBooks = append(UserDetails.IssuedBooks, cartBooks...)

	for _, book := range UserDetails.IssuedBooks {
		RemovefromCart(Userid, book.ID)
		book.Quantity = book.Quantity - 1
		db.Save(&book)
	}
	db.Save(&UserDetails)
	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}
	sqlDB.Close()
	return true, nil
}
func GetIssuedBooks(Userid int) ([]Book, error) {
	db, err := database.DbGormConnect()
	if err != nil {
		return nil, err
	}

	var IssuedBooks []Book
	if err := db.Joins("JOIN user_issued_books ON user_issued_books.book_id = books.id").
		Where("user_issued_books.user_id = ?", Userid).
		Find(&IssuedBooks).Error; err != nil {
		return nil, err
	}

	return IssuedBooks, nil
}
func ReturnBookByUser(UserId, BookId int) (bool, error) {
	db, err := database.DbGormConnect()
	if err != nil {
		return false, err
	}
	query := fmt.Sprintf("Delete from user_issued_books where user_id=%d and book_id=%d ;", UserId, BookId)

	err = db.Exec(query).Error
	if err != nil {
		return false, err
	}
	BookDetails := Book{}
	BookDetails, err = GetBookById(BookId)
	if err != nil {
		return false, err
	}
	BookDetails.Quantity = BookDetails.Quantity + 1
	db.Save(&BookDetails)
	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}
	sqlDB.Close()
	return true, nil

}
func GetUserLikedBooks(userId int) ([]Book, error) {
	db, err := database.DbGormConnect()
	if err != nil {
		return nil, err
	}

	var LikedBooks []Book
	if err := db.Joins("JOIN user_liked_books ON user_liked_books.book_id = books.id").
		Where("user_liked_books.user_id = ?", userId).
		Find(&LikedBooks).Error; err != nil {
		return nil, err
	}
	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}
	sqlDB.Close()

	return LikedBooks, nil
}
func GetVotesByBook(BookId int) (int, error) {
	db, err := database.DbGormConnect()
	if err != nil {
		return 0, err
	}
	var votes int
	err = db.Select("Count(*)").Where("book_id=" + strconv.Itoa(BookId)).Table("user_liked_books").Find(&votes).Error
	if err != nil {
		return 0, err
	}
	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}
	sqlDB.Close()
	return votes, nil
}
func GetUsers() ([]User, error) {
	db, err := database.DbGormConnect()
	if err != nil {
		return nil, err
	}
	var Users []User
	db.Find(&Users)
	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}
	sqlDB.Close()
	return Users, nil
}
