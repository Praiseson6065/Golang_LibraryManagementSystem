package handlers

import (
	"fmt"
	"log"
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
		err := rows.Scan(&book.BookId, &book.BookName, &book.ISBN, &book.Pages, &book.Publisher, &book.Author, &book.Taglines, &book.InsertedTime, &book.Votes)
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
		err := rows.Scan(&book.BookId, &book.BookName, &book.ISBN, &book.Pages, &book.Publisher, &book.Author, &book.Taglines, &book.InsertedTime, &book.Votes)
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
