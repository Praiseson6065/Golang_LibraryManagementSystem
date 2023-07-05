package handlers

import (
	"encoding/base64"
	"errors"
	"fmt"
	"strconv"

	"github.com/Praiseson6065/Golang_LibraryManagementSystem/database"
	"github.com/Praiseson6065/Golang_LibraryManagementSystem/middlewares"
	"github.com/Praiseson6065/Golang_LibraryManagementSystem/models"
	"github.com/gofiber/fiber/v2"
)

func AddBooksPost(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")
	claims, _ := middlewares.CookieGetData(cookie, c)

	if claims["usertype"] == "admin" {
		Book := new(models.Book)
		if err := c.BodyParser(Book); err != nil {

			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		file, err := c.FormFile("ImgPath")
		if err != nil {
			return err
		}
		Book.BookCode = base64.StdEncoding.EncodeToString([]byte(Book.BookName))
		path := "./static/img/books/" + Book.BookCode + ".jpg"
		Book.ImgPath = Book.BookCode + ".jpg"

		err = c.SaveFile(file, path)
		if err != nil {
			fmt.Println(err)
			return err
		}

		db, err := database.DbGormConnect()
		if err != nil {
			fmt.Println(err)
			return err
		}

		db.AutoMigrate(&models.Book{})

		db.Create(&Book)
		sqlDB, err := db.DB()
		if err != nil {
			panic(err)
		}
		sqlDB.Close()

		return c.JSON(true)

	} else {
		return c.JSON("unauthorized")
	}

}

func GetBooks(c *fiber.Ctx) error {

	Books, err := models.GetBooks()
	if err != nil {
		return err
	}
	return c.JSON(fiber.Map{
		"Books": Books,
	})

}

func SearchBooks(c *fiber.Ctx) error {
	Search := new(models.SearchBook)
	c.BodyParser(Search)

	Books, err := models.SearchBooks(Search.SearchValue, Search.SearchColumn)
	if err != nil {
		return err
	}
	return c.JSON(fiber.Map{
		"Books": Books,
	})

}
func GetBook(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return err
	}
	var book models.Book
	book, err = models.GetBookById(id)
	if err != nil {
		return err
	}
	book.Votes, err = models.GetVotesByBook(id)
	if err != nil {
		return err
	}
	return c.JSON(book)
}
func GetBookByCode(c *fiber.Ctx) error {
	bc := c.Params("bc")
	var Book models.Book
	Book, err := models.GetBookByBookCode(bc)
	if err != nil {
		return err
	}
	Book.Votes, err = models.GetVotesByBook(Book.ID)
	if err != nil {
		return err
	}
	return c.JSON(Book)
}

func UpdateBook(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid book ID",
		})
	}
	db, err := database.DbConnect()
	if err != nil {
		return err
	}
	Book := new(models.Book)
	c.BodyParser(Book)

	file, err := c.FormFile("ImgPath")
	if err != nil {
		if err == errors.New("there is no uploaded file associated with the given key") {
			fmt.Println("err")
		}
	}
	if file != nil {
		path := "./static/img/books/" + Book.BookCode + ".jpg"
		Book.ImgPath = Book.BookCode + ".jpg"

		err = c.SaveFile(file, path)
		if err != nil {
			fmt.Println("file")
			return err
		}
	}

	defer db.Close()
	updateQuery := "UPDATE books SET"
	updateArgs := []interface{}{}
	argCount := 1
	if Book.BookName != "" {
		updateQuery += fmt.Sprintf(" book_name=$%d,", argCount)
		updateArgs = append(updateArgs, Book.BookName)
		argCount++
	}

	if Book.ISBN != "" {
		updateQuery += fmt.Sprintf(" isbn=$%d,", argCount)
		updateArgs = append(updateArgs, Book.ISBN)
		argCount++
	}

	if Book.Pages != 0 {
		updateQuery += fmt.Sprintf(" pages=$%d,", argCount)
		updateArgs = append(updateArgs, Book.Pages)
		argCount++
	}

	if Book.Publisher != "" {
		updateQuery += fmt.Sprintf(" publisher=$%d,", argCount)
		updateArgs = append(updateArgs, Book.Publisher)
		argCount++
	}

	if Book.Author != "" {
		updateQuery += fmt.Sprintf(" author=$%d,", argCount)
		updateArgs = append(updateArgs, Book.Author)
		argCount++
	}

	if Book.Taglines != "" {
		updateQuery += fmt.Sprintf(" taglines=$%d,", argCount)
		updateArgs = append(updateArgs, Book.Taglines)
		argCount++
	}
	if file != nil {
		updateQuery += fmt.Sprintf(" img_path=$%d,", argCount)
		updateArgs = append(updateArgs, Book.ImgPath)
		argCount++
	}
	if Book.Quantity != 0 {
		updateQuery += fmt.Sprintf(" quantity = quantity+ $%d,", argCount)
		updateArgs = append(updateArgs, Book.Quantity)
		argCount++
	}
	updateQuery = updateQuery[:len(updateQuery)-1]
	updateQuery += fmt.Sprintf(" WHERE id=$%d", argCount)
	updateArgs = append(updateArgs, id)
	_, err = db.Exec(updateQuery, updateArgs...)
	defer db.Close()
	if err != nil {
		return err
	}
	return c.JSON(true)

}
func AddtoCart(c *fiber.Ctx) error {
	db, err := database.DbGormConnect()
	if err != nil {
		return err
	}
	bookid, err := strconv.Atoi(c.Params("bookid"))
	if err != nil {
		return err
	}
	userid, err := strconv.Atoi(c.Params("userid"))
	if err != nil {
		return err
	}
	UserDetails, err := models.GetUser(userid)
	if err != nil {
		return err
	}
	BookDetails, err := models.GetBookById(bookid)
	if err != nil {
		return err
	}
	UserDetails.CartBooks = append(UserDetails.CartBooks, BookDetails)
	db.Save(&UserDetails)
	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}
	sqlDB.Close()
	return c.JSON(fiber.Map{
		"msg": "AddedToCart",
	})
}

func RemoveFromCart(c *fiber.Ctx) error {
	bookId, err := strconv.Atoi(c.Params("bookid"))
	if err != nil {
		return err
	}

	userid, err := strconv.Atoi(c.Params("userid"))
	if err != nil {
		return err
	}

	_, err = models.RemovefromCart(userid, bookId)
	if err != nil {
		return err
	}

	return c.JSON(true)

}
func CheckOutFromCart(c *fiber.Ctx) error {
	Userid, err := strconv.Atoi(c.Params("userid"))
	if err != nil {
		return err
	}
	value, err := models.IssueBooks(Userid)
	if err != nil {
		return err
	}
	return c.JSON(value)

}
func ReturnBook(c *fiber.Ctx) error {
	bookId, err := strconv.Atoi(c.Params("bookid"))
	if err != nil {
		return err
	}

	userid, err := strconv.Atoi(c.Params("userid"))
	if err != nil {
		return err
	}
	value, err := models.ReturnBookByUser(userid, bookId)
	if err != nil {
		return err
	}
	return c.JSON(value)
}
func ChkBookCart(c *fiber.Ctx) error {
	bookId, err := strconv.Atoi(c.Params("bookid"))
	if err != nil {
		return err
	}

	userid, err := strconv.Atoi(c.Params("userid"))
	if err != nil {
		return err
	}
	var CartBooks []models.Book
	CartBooks, err = models.GetCartBooksByUserID(userid)
	if err != nil {
		return err
	}

	var book models.Book
	for _, book = range CartBooks {
		if book.ID == bookId {
			return c.JSON(true)
		}
	}
	return c.JSON(false)
}
func UserIssuedBooks(c *fiber.Ctx) error {
	userid, err := strconv.Atoi(c.Params("userid"))
	if err != nil {
		return err
	}
	var books []models.Book
	books, err = models.GetIssuedBooks(userid)
	if err != nil {
		return err
	}

	return c.JSON(books)
}
func UserApprovedBooks(c *fiber.Ctx) error {
	userid, err := strconv.Atoi(c.Params("userid"))
	if err != nil {
		return err
	}
	var ApprovBooks []models.Book
	ApprovBooks, err = models.GetUserApprovedBooks(userid)
	if err != nil {
		return err
	}
	return c.JSON(ApprovBooks)

}

func LikeBook(c *fiber.Ctx) error {
	db, err := database.DbGormConnect()
	if err != nil {
		return err
	}

	bookId, err := strconv.Atoi(c.Params("bookid"))
	if err != nil {
		return err
	}

	userid, err := strconv.Atoi(c.Params("userid"))
	if err != nil {
		return err
	}

	IsLiked := false
	if !IsLiked {
		var UserLikedBooks []models.Book
		UserLikedBooks, err = models.GetUserLikedBooks(userid)

		if err != nil {
			return err
		}
		var LikedBook models.Book
		for _, LikedBook = range UserLikedBooks {
			if LikedBook.ID == bookId {
				Query := fmt.Sprintf("Delete from user_liked_books where book_id=%d and user_id=%d ;", bookId, userid)
				err = db.Exec(Query).Error
				LikedBook.Votes = LikedBook.Votes - 1
				db.Save(&LikedBook)
				if err != nil {
					return c.JSON(err)
				}
				return c.JSON(false)
			}

		}

	}

	UserDetails := models.User{}

	UserDetails, err = models.GetUser(userid)
	if err != nil {
		return err
	}
	var LikedBook models.Book
	LikedBook, err = models.GetBookById(bookId)

	if err != nil {
		return err
	}
	LikedBook.Votes = LikedBook.Votes + 1
	db.Save(&LikedBook)
	UserDetails.LikedBooks = append(UserDetails.LikedBooks, LikedBook)
	db.Save(&UserDetails)

	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}
	sqlDB.Close()
	return c.JSON(true)

}
func IsLiked(c *fiber.Ctx) error {
	bookId, err := strconv.Atoi(c.Params("bookid"))
	if err != nil {
		return err
	}

	userid, err := strconv.Atoi(c.Params("userid"))
	if err != nil {
		return err
	}
	var LikedBooks []models.Book
	LikedBooks, err = models.GetUserLikedBooks(userid)
	if err != nil {
		return err
	}
	var book models.Book
	for _, book = range LikedBooks {
		if book.ID == bookId {
			return c.JSON(true)
		}
	}
	return c.JSON(false)

}
func IsBookIssued(c *fiber.Ctx) error {
	userid, err := strconv.Atoi(c.Params("userid"))
	if err != nil {
		return err
	}
	bookid, err := strconv.Atoi(c.Params("bookid"))
	if err != nil {
		return err
	}
	var IssuedBooksByUser []models.Book
	IssuedBooksByUser, err = models.GetIssuedBooks(userid)
	if err != nil {
		return err
	}
	var book models.Book
	for _, book = range IssuedBooksByUser {
		if book.ID == bookid {
			return c.JSON(true)
		}
	}
	return c.JSON(false)
}
