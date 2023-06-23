package handlers

import (
	"github.com/Praiseson6065/Golang_LibraryManagementSystem/models"
	"github.com/gofiber/fiber/v2"
)

func Userslist(c *fiber.Ctx) error{
	var Users []models.User
	Users,err:=models.GetUsers()
	if err!=nil{
		return err
	}
	return c.JSON(Users)
}