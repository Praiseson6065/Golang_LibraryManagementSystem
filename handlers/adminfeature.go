package handlers

import (
	"fmt"

	"github.com/Praiseson6065/Golang_LibraryManagementSystem/database"
	"github.com/Praiseson6065/Golang_LibraryManagementSystem/models"
	"github.com/gofiber/fiber/v2"
)

func AddBooksGet(c *fiber.Ctx) error {
	return c.Render("addbooks", map[string]interface{}{})

}
func AddBooksPost(c *fiber.Ctx) error {
	Book := new(models.Book)
	if err := c.BodyParser(Book); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	db, err := database.DbConnect()
	if err != nil {
		fmt.Println(err)
		return err
	}

	stmt, err := db.Prepare("INSERT INTO books (bookname,isbn,pages,publisher,author,quantity,taglines) values($1,$2,$3,$4,$5,$6)")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(stmt, Book.BookName, Book.ISBN, Book.Pages, Book.Publisher, Book.Author, Book.Quantity, Book.Taglines)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer stmt.Close()
	return c.Redirect("/BooksAdded")

}
