package handlers

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/Praiseson6065/Golang_LibraryManagementSystem/database"
	"github.com/Praiseson6065/Golang_LibraryManagementSystem/models"
	"github.com/gofiber/fiber/v2"
)

func GetBooks(c *fiber.Ctx) error {
	db, err := database.DbConnect()
	if err != nil {
		fmt.Println(err)
		return err
	}
	rows, err := db.Query("SELECT * FROM books")
	if err != nil {
		log.Fatal(err)
	}
	var Books []models.Book
	for rows.Next() {
		var book models.Book
		err := rows.Scan(&book.BookId, &book.BookName, &book.ISBN, &book.Pages, &book.Publisher, &book.Author, &book.Taglines, &book.InsertedTime, &book.Votes, &book.ImgPath)
		if err != nil {
			log.Fatal(err)
		}
		Books = append(Books, book)
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}
	db.Close()
	return c.JSON(fiber.Map{
		"Books": Books,
	})

}

func SearchBooks(c *fiber.Ctx) error {
	Search := new(models.SearchBook)
	c.BodyParser(Search)
	db, err := database.DbConnect()
	if err != nil {
		fmt.Println(err)
		return err
	}
	rows, err := db.Query("SELECT * FROM books WHERE LOWER(" + Search.SearchColumn + ") LIKE '%" + strings.ToLower(Search.SearchValue) + "%' ;")
	if err != nil {
		log.Fatal(err)
	}
	var Books []models.Book
	for rows.Next() {
		var book models.Book
		err := rows.Scan(&book.BookId, &book.BookName, &book.ISBN, &book.Pages, &book.Publisher, &book.Author, &book.Taglines, &book.InsertedTime, &book.Votes, &book.ImgPath)
		if err != nil {
			log.Fatal(err)
		}
		Books = append(Books, book)
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}
	db.Close()
	return c.JSON(fiber.Map{
		"Books": Books,
	})

}
func GetBook(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.SendStatus(500)
	}
	db, err := database.DbConnect()
	if err != nil {
		fmt.Println(err)
		return err
	}
	Book := new(models.Book)
	Query := "Select * From books where bookid=$1 ;"

	result, err := db.Query(Query, id)
	if err != nil {
		return err
	}
	if !result.Next() {
		return c.SendStatus(405)
	}
	err = result.Scan(&Book.BookId, &Book.BookName, &Book.ISBN, &Book.Pages, &Book.Publisher, &Book.Author, &Book.Taglines, &Book.InsertedTime, &Book.Votes, &Book.ImgPath)
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
		path := "./static/img/" + Book.ISBN
		Book.ImgPath = Book.ISBN

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
		updateQuery += fmt.Sprintf(" bookname=$%d,", argCount)
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
		updateQuery += fmt.Sprintf(" img=$%d,", argCount)
		updateArgs = append(updateArgs, Book.ImgPath)
		argCount++
	}

	// Remove the trailing comma and complete the query
	updateQuery = updateQuery[:len(updateQuery)-1]
	updateQuery += fmt.Sprintf(" WHERE bookid=$%d", argCount)
	updateArgs = append(updateArgs, id)
	_, err = db.Exec(updateQuery, updateArgs...)
	if err != nil {
		return err
	}
	return c.JSON(fiber.Map{
		"msg": "Successfully Updated",
	})

}
