package handlers

import (
	"github.com/Praiseson6065/Golang_LibraryManagementSystem/middlewares"
	"github.com/Praiseson6065/Golang_LibraryManagementSystem/database"
	"github.com/Praiseson6065/Golang_LibraryManagementSystem/models"
	"github.com/gofiber/fiber/v2"
	"encoding/base64"
)

func Userslist(c *fiber.Ctx) error{
	var Users []models.User
	Users,err:=models.GetUsers()
	if err!=nil{
		return err
	}
	return c.JSON(Users)
}
func AddAdmin(c *fiber.Ctx) error{
	var user models.User
	err:= c.BodyParser(&user)
	if err!=nil{
		return c.JSON(err)
	}
	db,err:=database.DbGormConnect()
	if err!=nil{
		return c.JSON(err)
	}
	
	db.AutoMigrate(&models.User{})
	user.Usertype="admin"
	hashpassword, err := middlewares.HashPassword(user.Password)
	if err != nil {
		return err
	}
	user.Password=hashpassword
	user.UserId = base64.StdEncoding.EncodeToString([]byte(user.Name))
	err=db.Create(&user).Error
	if err!=nil{
		return c.JSON(err)
	}
	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}
	sqlDB.Close()
	return c.JSON(true)

}
func ReqBook(c *fiber.Ctx) error{
	db,err:= database.DbGormConnect()
	if err!=nil{
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