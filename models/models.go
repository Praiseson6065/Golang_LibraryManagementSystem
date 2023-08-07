package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"github.com/Praiseson6065/Golang_LibraryManagementSystem/database"
	"gorm.io/gorm"
)

type Pi struct {
	Id           string `json:"id"`
	ClientSecret string `json:"client_secret"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token string `json:"token"`
}
type Log struct {
	UserId    int
	BookId    int
	Operation string
}
type UserBookLog struct {
	gorm.Model
	Book      int
	Operation string
}
type CartPurchaseBooks struct {
	gorm.Model
	Book     int
	Quantity int
}
type PurchaseCart struct {
	Book            Book
	PurchaseDetails CartPurchaseBooks
}
type User struct {
	gorm.Model
	ID                int                 `json:"Id" gorm:"auto_increment:true;primary_key;unique"`
	UserId            string              `json:"UserId"`
	Name              string              `json:"Name"`
	Email             string              `json:"Email" gorm:"unique"`
	Password          string              `json:"Password"`
	Usertype          string              `json:"Usertype"`
	CartBooks         []Book              `json:"CartBooks" gorm:"many2many:user_cart_books;"`
	LikedBooks        []Book              `json:"LikedBooks" gorm:"many2many:user_liked_books;"`
	IssuedBooks       []Book              `json:"IssuedBooks" gorm:"many2many:user_issued_books;"`
	ApprovedBooks     []Book              `json:"ApprovedBooks" gorm:"many2many:user_approved_books;"`
	UserLog           []UserBookLog       `json:"UserBookLog" gorm:"many2many:user_log_books;"`
	CartPurchaseBooks []CartPurchaseBooks `json:"CartPurchase" gorm:"many2many:user_cartpurchase_books"`
	PurchasedBooks    []CartPurchaseBooks `json:"PurchasedBooks" gorm:"many2many:user_purchased_books"`
}

type UserBookDetails struct {
	Approve bool `json:"Approve"`
	Cart    bool `json:"Cart"`
	Issued  bool `json:"Issued"`
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
type BookReviews struct {
	gorm.Model
	UserId   int    `json:"UserId"`
	UserName string `json:"UserName"`
	BookId   int    `json:"BookId"`
	Review   string `json:"Review"`
}
type Book struct {
	gorm.Model
	ID           int     `json:"BookId" gorm:"auto_increment:true;primary_key;unique"`
	BookCode     string  `json:"BookCode" gorm:"unique"`
	BookName     string  `json:"BookName"`
	ISBN         string  `json:"ISBN" gorm:"unique"`
	Pages        int     `json:"Pages"`
	Publisher    string  `json:"Publisher"`
	Quantity     int     `json:"Quantity"`
	Author       string  `json:"Author"`
	Price        float64 `json:"Price"`
	Taglines     string  `json:"Taglines"`
	InsertedTime string  `json:"InsertedTime"`
	Votes        int     `json:"votes"`
	ImgPath      string  `json:"ImgPath"`
}
type SearchBook struct {
	SearchValue  string `json:"SearchValue"`
	SearchColumn string `json:"SearchColumn"`
}
type ApprovBooks struct {
	Userid  int   `json:"UserId"`
	BookIds []int `json:"BookId"`
}

type GooglePayload struct {
	SUB           string `json:"sub"`
	Name          string `json:"name"`
	GivenName     string `json:"given_name"`
	FamilyName    string `json:"family_name"`
	Picture       string `json:"picture"`
	Email         string `json:"email"`
	EmailVerified bool   `json:"email_verified"`
	Locale        string `json:"locale"`
}

func ConvertToken(accessToken string) (*GooglePayload, error) {

	resp, httpErr := http.Get(fmt.Sprintf("https://www.googleapis.com/oauth2/v3/userinfo?access_token=%s", accessToken))
	if httpErr != nil {
		return nil, httpErr
	}

	defer resp.Body.Close()

	respBody, bodyErr := ioutil.ReadAll(resp.Body)
	if bodyErr != nil {
		return nil, bodyErr
	}

	var body map[string]interface{}
	if err := json.Unmarshal(respBody, &body); err != nil {
		return nil, err
	}

	if body["error"] != nil {
		return nil, errors.New("invalid token")
	}

	var data GooglePayload
	err := json.Unmarshal(respBody, &data)
	if err != nil {
		return nil, err
	}

	return &data, nil
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
	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}
	sqlDB.Close()
	return IssuedBooks, nil
}
func RemoveIssueBook(UserId, BookId int) (bool, error) {
	db, err := database.DbGormConnect()
	if err != nil {
		return false, err
	}
	query := fmt.Sprintf("Delete from user_issued_books where user_id=%d and book_id=%d ;", UserId, BookId)

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
func RemoveIssueBookByUser(UserId, BookId int) (bool, error) {
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
func LogOperation(userid, bookid int, opr string) bool {
	db, err := database.DbGormConnect()
	if err != nil {
		fmt.Println(err)
		return false
	}
	db.AutoMigrate(&Log{})
	LogOpr := Log{
		UserId:    userid,
		BookId:    bookid,
		Operation: opr,
	}
	err = db.Create(&LogOpr).Error
	if err != nil {
		fmt.Println(err)
		return false
	}
	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}
	sqlDB.Close()
	return true
}
func ReturnBookByUser(UserId, BookId int) (bool, error) {
	db, err := database.DbGormConnect()
	if err != nil {
		return false, err
	}
	query := fmt.Sprintf("Delete from user_approved_books where user_id=%d and book_id=%d ;", UserId, BookId)

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
func ApprovBook(userid int, BookId []int) (bool, error) {
	db, err := database.DbGormConnect()
	if err != nil {
		return false, err
	}
	var ApprovBooks []Book
	for _, i := range BookId {
		RemoveIssueBook(userid, i)
		book, err := GetBookById(i)
		if err != nil {
			return false, err
		}
		ApprovBooks = append(ApprovBooks, book)

	}

	UserDetails, err := GetUser(userid)
	if err != nil {
		return false, err
	}
	UserDetails.ApprovedBooks = append(UserDetails.ApprovedBooks, ApprovBooks...)
	db.Save(&UserDetails)
	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}
	sqlDB.Close()

	return true, nil

}
func GetUserApprovedBooks(userid int) ([]Book, error) {
	db, err := database.DbGormConnect()
	if err != nil {
		return nil, err
	}
	var AppRovBooks []Book
	if err := db.Joins("JOIN user_approved_books ON user_approved_books.book_id = books.id").
		Where("user_approved_books.user_id = ?", userid).
		Find(&AppRovBooks).Error; err != nil {
		return nil, err
	}
	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}
	sqlDB.Close()
	return AppRovBooks, nil

}
func GetReviewByUserBookId(UserId, BookId int) (BookReviews, error) {
	db, err := database.DbGormConnect()
	if err != nil {
		return BookReviews{}, err
	}

	var bkRvw BookReviews

	err = db.Where("user_id=? and book_id=?", UserId, BookId).Find(&bkRvw).Error
	if err != nil {
		return BookReviews{}, err
	}
	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}
	sqlDB.Close()
	return bkRvw, nil
}
func GetReviewsByBookId(BookId int) ([]BookReviews, error) {
	db, err := database.DbGormConnect()
	if err != nil {
		return nil, err
	}
	var BookReviews []BookReviews
	err = db.Where("book_id=?", BookId).Find(&BookReviews).Error
	if err != nil {
		return nil, err
	}
	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}
	sqlDB.Close()
	return BookReviews, nil
}
func UserPurchaseCart(userid int) ([]CartPurchaseBooks, error) {
	db, err := database.DbGormConnect()
	if err != nil {
		return nil, err
	}
	var cartPurchaseBooks []CartPurchaseBooks
	if err := db.Joins("JOIN user_cartpurchase_books ON user_cartpurchase_books.cart_purchase_books_id  = cart_purchase_books.id").
		Where("user_cartpurchase_books.user_id = ?", userid).
		Find(&cartPurchaseBooks).Error; err != nil {
		return nil, err
	}
	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}
	sqlDB.Close()
	return cartPurchaseBooks, nil

}
func AddtoPurchaseCart(userid, bookid, quantity int) (bool, error) {
	db, err := database.DbGormConnect()
	if err != nil {
		return false, err
	}
	UserD, err := GetUser(userid)
	if err != nil {
		return false, err
	}
	Book, err := GetBookById(bookid)
	if err != nil {
		return false, err
	}
	userCartP, err := UserPurchaseCart(userid)
	if err != nil {
		return false, err
	}
	ans := false
	for _, i := range userCartP {
		if i.Book == bookid {

			ans = true
			i.Quantity = i.Quantity + quantity
			if i.Quantity > 0 {
				db.Save(&i)
			} else {
				RemovefromPurchaseCart(userid, bookid)
			}

			break
		}
	}
	if !ans {
		CPurchaseBook := CartPurchaseBooks{
			Book:     Book.ID,
			Quantity: quantity,
		}

		UserD.CartPurchaseBooks = append(UserD.CartPurchaseBooks, CPurchaseBook)
		db.Save(&UserD)

	}

	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}
	sqlDB.Close()

	return true, nil
}
func RemovefromPurchaseCart(userid, bookid int) (bool, error) {
	db, err := database.DbGormConnect()
	if err != nil {
		return false, err
	}
	userCartP, err := UserPurchaseCart(userid)
	if err != nil {
		return false, err
	}
	for _, i := range userCartP {
		if i.Book == bookid {
			db.Delete(&i)
			query := fmt.Sprintf("delete from user_cartpurchase_books where cart_purchase_books_id=%d;", i.ID)
			err = db.Exec(query).Error
			if err != nil {
				return false, err
			}
			break
		}
	}
	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}
	sqlDB.Close()

	return true, nil
}
func PurchaseBooks(userid int) (bool, error) {
	db, err := database.DbGormConnect()
	if err != nil {
		return false, err
	}

	UserD, err := GetUser(userid)
	if err != nil {
		return false, err
	}
	userCartP, err := UserPurchaseCart(userid)
	if err != nil {
		return false, err
	}
	for _, i := range userCartP {
		book, err := GetBookById(i.Book)
		if err != nil {
			return false, err
		}
		book.Quantity = book.Quantity - i.Quantity
		db.Save(&book)
	}
	UserD.PurchasedBooks = append(UserD.PurchasedBooks, userCartP...)
	db.Save(&UserD)
	query := fmt.Sprintf("delete from user_cartpurchase_books where user_id=%d;", userid)
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
func GetPurchaseCartbyId(userid int) ([]PurchaseCart, error) {

	var cartPurchaseBook []CartPurchaseBooks
	cartPurchaseBook, err := UserPurchaseCart(userid)
	if err != nil {
		return nil, err
	}
	var PurchaseCrt []PurchaseCart
	for _, i := range cartPurchaseBook {
		book, err := GetBookById(i.Book)
		if err != nil {
			return nil, err
		}
		var Pc PurchaseCart
		Pc.Book = book
		Pc.PurchaseDetails = i
		PurchaseCrt = append(PurchaseCrt, Pc)
	}
	return PurchaseCrt, nil

}
