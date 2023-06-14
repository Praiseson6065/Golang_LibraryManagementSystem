package handlers

import (
	"fmt"

	"github.com/Praiseson6065/Golang_LibraryManagementSystem/database"
	"github.com/Praiseson6065/Golang_LibraryManagementSystem/middlewares"
	"github.com/Praiseson6065/Golang_LibraryManagementSystem/models"
	"github.com/gofiber/fiber/v2"
)

func AddBooksPost(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")
	claims, _ := middlewares.CookieGetData(cookie, c)
	fmt.Println(claims)
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

		path := "./static/img/" + Book.ISBN
		Book.ImgPath = Book.ISBN

		err = c.SaveFile(file, path)
		if err != nil {
			fmt.Println("file")
			return err
		}

		db, err := database.DbConnect()
		if err != nil {
			fmt.Println(err)
			return err
		}

		stmt, err := db.Prepare("INSERT INTO books (bookname,isbn,pages,publisher,author,taglines,img) values($1,$2,$3,$4,$5,$6,$7)")
		if err != nil {
			return err
		}
		_, err = stmt.Exec(Book.BookName, Book.ISBN, Book.Pages, Book.Publisher, Book.Author, Book.Taglines, Book.ImgPath)
		if err != nil {
			fmt.Println(err)
			return err
		}
		defer stmt.Close()
		return c.JSON(fiber.Map{
			"Message": "Book Added Succesfully",
		})

	} else {
		return c.JSON(fiber.Map{
			"msg": "unauthorized",
		})
	}

}
