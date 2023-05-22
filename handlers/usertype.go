package handlers

import (
	
	"log"

	"github.com/Praiseson6065/Golang_LibraryManagementSystem/database"
	"github.com/Praiseson6065/Golang_LibraryManagementSystem/middlewares"
	"github.com/Praiseson6065/Golang_LibraryManagementSystem/models"
	"github.com/gofiber/fiber/v2"
)

func AdminPage(c *fiber.Ctx) error {
	db, err := database.DbConnect()

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

	return c.Render("admin", fiber.Map{
		"Data": data,
	})
}
func ProfilePage(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")
	claims, err := middlewares.CookieGetData(cookie, c)
	
	

	if err != nil {
		return c.Render("profile", map[string]interface{}{
			"name":       claims["name"],
			"email":      claims["email"],
			"hideEle":    "hide-data",
			"user-st":    "/logout",
			"userstatus": "Logout",
		})
	} else {
		return c.Render("profile", map[string]interface{}{
			"msg":        "Not Logged In.",
			"nologin":    "hide-data",
			"user-st":    "/login",
			"userstatus": "Login",
			"hideEle":    "",
		})
	}

}
