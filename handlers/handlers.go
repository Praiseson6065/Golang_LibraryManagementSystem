package handlers

import (
	"fmt"

	"encoding/base64"
	"time"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"

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
	t, err := token.SignedString([]byte(config.EnvConfigs.SecretKey))
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
		// fmt.Println(err)
		return c.JSON(err)
	}

	data := new(models.User)
	c.BodyParser(data)
	hashpassword, err := middlewares.HashPassword(data.Password)
	if err != nil {
		return c.JSON(err)
	}
	user := models.User{
		Email:    data.Email,
		UserId:   base64.StdEncoding.EncodeToString([]byte(data.Name)),
		Password: hashpassword,
		Name:     data.Name,
		Usertype: "user",
	}

	db.AutoMigrate(&models.User{})

	err = db.Create(&user).Error
	if err != nil {
		return c.JSON(err)
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

func GoogleAuthLogin(c *fiber.Ctx) error {
	conf := &oauth2.Config{
		ClientID:     config.EnvConfigs.G_CLIENT_ID,
		ClientSecret: config.EnvConfigs.G_CLIENT_SECRET,
		RedirectURL:  config.EnvConfigs.G_REDIRECT,
		Endpoint:     google.Endpoint,
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email", "https://www.googleapis.com/auth/userinfo.profile"},
	}
	URL := conf.AuthCodeURL("golanglibrary")
	return c.Redirect(URL)
}
func GoogleCallBack(c *fiber.Ctx) error {
	conf := &oauth2.Config{
		ClientID:     config.EnvConfigs.G_CLIENT_ID,
		ClientSecret: config.EnvConfigs.G_CLIENT_SECRET,
		RedirectURL:  config.EnvConfigs.G_REDIRECT,
		Endpoint:     google.Endpoint,
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email", "https://www.googleapis.com/auth/userinfo.profile"},
	}
	code := c.Query("code")
	token, err := conf.Exchange(c.Context(), code)
	fmt.Println("Refesh :", token.RefreshToken, "Access :", token.AccessToken)
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	profile, err := models.ConvertToken(token.AccessToken)
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	_, err = repository.FindByGoogleAcc(profile.Email, profile.SUB)
	if err != nil {
		if err.Error() == "user not found" {

			db, err := database.DbGormConnect()
			if err != nil {
				return c.JSON(err)
			}

			user := models.User{
				Email:    profile.Email,
				UserId:   base64.StdEncoding.EncodeToString([]byte(profile.GivenName)),
				Password: profile.SUB,
				Name:     profile.GivenName,
				Usertype: "user",
			}
			err = db.Create(&user).Error
			if err != nil {
				return c.JSON(err)
			}
		}

	}
	day := time.Hour * 24
	FindUserCredReg, err := repository.FindByGoogleAcc(profile.Email, profile.SUB)

	if err != nil {

		return c.JSON(err)
	}

	claims := jtoken.MapClaims{
		"ID":       FindUserCredReg.ID,
		"email":    FindUserCredReg.Email,
		"name":     FindUserCredReg.Name,
		"usertype": FindUserCredReg.Usertype,
		"exp":      time.Now().Add(day * 1).Unix(),
	}

	jwttoken := jtoken.NewWithClaims(jtoken.SigningMethodHS256, claims)

	t, err := jwttoken.SignedString([]byte(config.EnvConfigs.SecretKey))
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
	if FindUserCredReg.Usertype == "user" {
		return c.Redirect("/profile.html")
	} else if FindUserCredReg.Usertype == "admin" {
		return c.Redirect("/admin.html")
	} else {
		return c.JSON(fiber.Map{
			"msg": "Invalid",
		})
	}

}
