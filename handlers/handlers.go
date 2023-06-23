package handlers

import (
	"fmt"
	"strconv"

	"encoding/base64"
	"time"

	"github.com/Praiseson6065/Golang_LibraryManagementSystem/config"
	"github.com/Praiseson6065/Golang_LibraryManagementSystem/database"
	"github.com/Praiseson6065/Golang_LibraryManagementSystem/middlewares"
	"github.com/Praiseson6065/Golang_LibraryManagementSystem/models"
	"github.com/Praiseson6065/Golang_LibraryManagementSystem/repository"
	"github.com/gofiber/fiber/v2"
	jtoken "github.com/golang-jwt/jwt/v4"
)

// Login route
func Login(c *fiber.Ctx) error {

	loginRequest := new(models.LoginRequest)
	if err := c.BodyParser(loginRequest); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	user, err := repository.FindByCredentials(loginRequest.Email, loginRequest.Password)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	day := time.Hour * 24

	claims := jtoken.MapClaims{
		"ID":       user.ID,
		"email":    user.Email,
		"name":     user.Name,
		"usertype": user.Usertype,
		"exp":      time.Now().Add(day * 1).Unix(),
	}

	token := jtoken.NewWithClaims(jtoken.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(config.Secret))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	c.Cookie(&fiber.Cookie{
		Name:    "jwt",
		Value:   t,
		Expires: time.Now().Add(24 * time.Hour),
	})
	if user.Usertype == "user" {
		return c.Redirect("/profile.html")
	} else if user.Usertype == "admin" {
		return c.Redirect("/admin.html")
	} else {
		return c.JSON(fiber.Map{
			"msg": "Invalid",
		})
	}

}

// Protected route
func Protected(c *fiber.Ctx) error {
	// Get the user from the context and return it
	user := c.Locals("user").(*jtoken.Token)
	claims := user.Claims.(jtoken.MapClaims)
	email := claims["email"].(string)
	Name := claims["name"].(string)
	fmt.Println("claims:", claims)
	return c.SendString("Welcome 👋" + email + " " + Name)
}

//Login Page
func Loginpage(c *fiber.Ctx) error {

	cookie := c.Cookies("jwt")
	claims, _ := middlewares.CookieGetData(cookie, c)
	if claims["usertype"] == "user" {
		return c.Redirect("/profile.html")
	} else if claims["usertype"] == "admin" {
		return c.Redirect("/admin.html")
	} else {
		return c.Redirect(("/login.html"))
	}

}

//register

func RegisterPost(c *fiber.Ctx) error {
	db, err := database.DbGormConnect()

	if err != nil {
		fmt.Println(err)
		return err
	}
	db.AutoMigrate(&models.User{})
	data := new(models.User)
	c.BodyParser(data)
	hashpassword, err := middlewares.HashPassword(data.Password)
	if err != nil {
		return err
	}
	user := models.User{
		Email:    data.Email,
		UserId:   base64.StdEncoding.EncodeToString([]byte(data.Name)),
		Password: hashpassword,
		Name:     data.Name,
		Usertype: "user",
	}
	err = db.Create(&user).Error
	if err != nil {
		return err
	}
	return c.JSON(true)
}
func RegisterSuccessful(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"msg": "Registration Succesful",
	})

}
func Logout(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")
	claims, err := middlewares.CookieGetData(cookie, c)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Not Logged In",
		})
	}

	if claims["name"] != "nologin" {
		c.ClearCookie()
		return c.JSON(fiber.Map{
			"logoutmsg": claims["name"],
		})
	} else {
		return c.Render("logout", map[string]interface{}{
			"title":     "Logout",
			"logoutmsg": "Not Logged In.",
		})
	}

}

func GetUserCart(c *fiber.Ctx) error {
	Userid, err := strconv.Atoi(c.Params("userid"))
	if err != nil {
		return err
	}
	var CBooks []models.Book
	CBooks, err = models.GetCartBooksByUserID(Userid)
	if err != nil {
		return err
	}
	return c.JSON(CBooks)
}

