package handlers

import (
	"log"

	"github.com/Praiseson6065/Golang_LibraryManagementSystem/database"
	"github.com/Praiseson6065/Golang_LibraryManagementSystem/middlewares"
	"github.com/Praiseson6065/Golang_LibraryManagementSystem/models"
	"github.com/gofiber/fiber/v2"
)

func AdminPage(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")
	claims, _ := middlewares.CookieGetData(cookie, c)
	if claims["usertype"] == "admin" {
		db, err := database.DbConnect()
		if err!=nil{
			return err
		}
		// Query the database for the data to display.
		rows, err := db.Query("SELECT * FROM logdb")
		if err != nil {
			log.Fatal(err)
		}

		var data []models.Logdb

		for rows.Next() {
			var d models.Logdb
			err := rows.Scan(&d.LogId, &d.UserId, &d.UserType, &d.Operation, &d.InsertTime, &d.UserName)
			if err != nil {
				log.Fatal(err)
			}
			data = append(data, d)

		}

		if err = rows.Err(); err != nil {
			log.Fatal(err)
		}
		db.Close()

		return c.JSON(fiber.Map{
			"Data": data,
		})

	} else {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "UnAuthorized",
		})
	}

}
func ProfilePage(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")
	claims, err := middlewares.CookieGetData(cookie, c)

	if err != nil {
		return c.JSON(fiber.Map{
			"msg": "Not Logged In",
		})
	} else {
		return c.JSON(fiber.Map{
			"msg":   "Logged In",
			"name":  claims["name"],
			"email": claims["email"],
		})
	}

}
