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

		return c.JSON(fiber.Map{
			"Message": "Book Added Succesfully",
		})

	} else {
		return c.JSON(fiber.Map{
			"msg": "unauthorized",
		})
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
	book := models.Book{}
	book, err = models.GetBookById(id)
	if err != nil {
		return err
	}
	return c.JSON(book)
}
func GetBookByCode(c *fiber.Ctx) error {
	bc := c.Params("bc")
	Book := models.Book{}
	Book, err := models.GetBookByBookCode(bc)
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
	if err != nil {
		return err
	}
	return c.JSON(fiber.Map{
		"msg": "Successfully Updated",
	})

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
	db, err := database.DbGormConnect()
	if err != nil {
		return err
	}

	if err != nil {
		return err
	}

	query := fmt.Sprintf("Delete from user_cart_books where user_id=%d and book_id=%d ;", userid, bookId)

	err = db.Exec(query).Error
	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{
		"msg": "Removed",
	})

}
