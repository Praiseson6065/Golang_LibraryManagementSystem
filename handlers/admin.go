package handlers

import (
	"encoding/base64"

	"github.com/Praiseson6065/Golang_LibraryManagementSystem/database"
	"github.com/Praiseson6065/Golang_LibraryManagementSystem/middlewares"
	"github.com/Praiseson6065/Golang_LibraryManagementSystem/models"
	"github.com/gofiber/fiber/v2"
)

func Userslist(c *fiber.Ctx) error {
	var Users []models.User
	Users, err := models.GetUsers()
	if err != nil {
		return err
	}
	return c.JSON(Users)
}
func AddAdmin(c *fiber.Ctx) error {
	var user models.User
	err := c.BodyParser(&user)
	if err != nil {
		return c.JSON(err)
	}
	db, err := database.DbGormConnect()
	if err != nil {
		return c.JSON(err)
	}

	db.AutoMigrate(&models.User{})
	user.Usertype = "admin"
	hashpassword, err := middlewares.HashPassword(user.Password)
	if err != nil {
		return err
	}
	user.Password = hashpassword
	user.UserId = base64.StdEncoding.EncodeToString([]byte(user.Name))
	err = db.Create(&user).Error
	if err != nil {
		return c.JSON(err)
	}
	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}
	sqlDB.Close()
	return c.JSON(true)

}
func ReqBook(c *fiber.Ctx) error {
	db, err := database.DbGormConnect()
	if err != nil {
		return c.JSON(err)
	}
	var ReqBooks []models.UserRequestedBooks
	db.Find(&ReqBooks)
	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}
	sqlDB.Close()
	return c.JSON(ReqBooks)
}
func ApprovBookByAdmin(c *fiber.Ctx) error {
	var ApprBooks []models.ApprovBooks
	c.BodyParser(&ApprBooks)
	for _, i := range ApprBooks {
		_, err := models.ApprovBook(i.Userid, i.BookIds)
		if err != nil {
			return c.JSON(err)
		}
	}
	return c.JSON(true)
}
func GetApprovalBooks(c *fiber.Ctx) error {

	type UserData struct {
		User          models.User
		CartBooks     []models.Book
		LikedBooks    []models.Book
		IssuedBooks   []models.Book
		ApprovedBooks []models.Book
	}
	var UsrData []UserData
	users, err := models.GetUsers()
	if err != nil {
		return c.JSON(err)
	}

	for _, i := range users {
		var userApprovedBooks []models.Book
		userApprovedBooks, err = models.GetUserApprovedBooks(i.ID)
		if err != nil {
			return c.JSON(err)
		}
		var Ud UserData
		Ud.User = i
		Ud.ApprovedBooks = userApprovedBooks

		var userIssuedBooks []models.Book
		userIssuedBooks, err = models.GetIssuedBooks(i.ID)
		if err != nil {
			return c.JSON(err)
		}
		Ud.IssuedBooks = userIssuedBooks
		UsrData = append(UsrData, Ud)

	}

	return c.JSON(UsrData)

}
