package handlers

import (
	"github.com/Praiseson6065/Golang_LibraryManagementSystem/middlewares"
	"github.com/gofiber/fiber/v2"
)

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
