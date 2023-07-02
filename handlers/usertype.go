package handlers

import (
	"strconv"

	"github.com/Praiseson6065/Golang_LibraryManagementSystem/database"
	"github.com/Praiseson6065/Golang_LibraryManagementSystem/middlewares"
	"github.com/Praiseson6065/Golang_LibraryManagementSystem/models"
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
func UserRequestedBooks(c *fiber.Ctx)error{
	db,err:=database.DbGormConnect()
	if err!=nil{
		return c.JSON(err)
	}
	var ReqBook models.UserRequestedBooks
	db.AutoMigrate(&models.UserRequestedBooks{})
	c.BodyParser(&ReqBook)
	ReqBook.RequestStatus=false
	err=db.Create(&ReqBook).Error
	if err!=nil{
		return c.JSON(err)
	}
	return c.JSON(true)
}
func RequestedBooks(c *fiber.Ctx) error{
	userId,err := strconv.Atoi(c.Params("userid"))
	if err!=nil{
		return c.JSON(err)
	}
	db,err:=database.DbGormConnect()
	if err!=nil{
		return c.JSON(err)
	}
	var ReqBooks []models.UserRequestedBooks

	db.Find(&ReqBooks).Where("user_id=?",userId)
	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}
	sqlDB.Close()
	return c.JSON(ReqBooks)
}