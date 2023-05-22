package handlers

import (
	
	"errors"
	"fmt"
	"time"

	"github.com/Praiseson6065/Golang_LibraryManagementSystem/config"
	"github.com/Praiseson6065/Golang_LibraryManagementSystem/database"
	"github.com/Praiseson6065/Golang_LibraryManagementSystem/middlewares"
	"github.com/Praiseson6065/Golang_LibraryManagementSystem/models"
	"github.com/Praiseson6065/Golang_LibraryManagementSystem/repository"
	"github.com/gofiber/fiber/v2"
	jtoken "github.com/golang-jwt/jwt/v4"
)

//home page
func HomePage(c *fiber.Ctx) error {
	return c.Render("home", map[string]interface{}{})
}

	
// Login route
func Login(c *fiber.Ctx) error {
	// Extract the credentials from the request body
	loginRequest := new(models.LoginRequest)
	if err := c.BodyParser(loginRequest); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	// Find the user by credentials

	user, err := repository.FindByCredentials(loginRequest.Email, loginRequest.Password)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	day := time.Hour * 24
	// Create the JWT claims, which includes the user ID and expiry time
	claims := jtoken.MapClaims{
		"ID":       user.ID,
		"email":    user.Email,
		"name":     user.Name,
		"usertype": user.Usertype,
		"exp":      time.Now().Add(day * 1).Unix(),
	}
	fmt.Printf("claims:%T/n",claims)
	// Create token
	token := jtoken.NewWithClaims(jtoken.SigningMethodHS256, claims)
	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte(config.Secret))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	db, err := database.DbConnect()
	stmt, err := db.Prepare("INSERT INTO logdb (userid,user_type,operation,userName) VALUES ($1, $2, $3, $4)")
	_, err = stmt.Exec(user.ID, user.Usertype, "login", user.Name)
	if err != nil {	
		fmt.Println(err)
		return err
	}
	defer stmt.Close()
	// c.Set("Authorization", "Bearer"+t)
	c.Cookie(&fiber.Cookie{
		Name:     "jwt",
		Value:    t,
		Expires:  time.Now().Add(24 * time.Hour),
		HTTPOnly: true,
		SameSite: "lax",
	})
	if user.Usertype == "user" {
		return c.Redirect("/profile")
	} else if user.Usertype == "admin" {
		return c.Redirect("/admin")
	} else {
		return nil
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
	return c.Render("login", map[string]interface{}{
		"title": "Login Page"})
}

//register
func Register(c *fiber.Ctx) error {
	return c.Render("register", map[string]interface{}{
		"title": "Register Page"})
}
func RegisterPost(c *fiber.Ctx) error {
	db, err := database.DbConnect()
	if err != nil {
		fmt.Println(err)
		return err
	}
	data := new(models.RegisterUser)
	c.BodyParser(data)
	// Insert the user data into the table.
	stmt, err := db.Prepare("INSERT INTO user_data (email, password, name,user_type) VALUES ($1, $2, $3,$4)")
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer stmt.Close()

	hashpassword, err := middlewares.HashPassword(data.Password)
	if err != nil {
		return nil
	}

	_, err = stmt.Exec(data.Email, hashpassword, data.Name, "user")
	if err != nil {
		fmt.Println(err)
		return err
	}

	query := `SELECT id, email, password, name,user_type FROM user_data WHERE email = $1 ;`

	// Execute the query
	result, err := db.Query(query, data.Email)
	if err != nil {
		return nil
	}
	// Check if the query returned any rows
	if !result.Next() {
		return errors.New("user not found")
	}
	user := models.User{}
	err = result.Scan(&user.ID, &user.Email, &user.Password, &user.Name, &user.Usertype)
	if err != nil {
		return nil
	}
	stmt, err = db.Prepare("INSERT INTO logdb (userid,user_type,operation,userName) VALUES ($1, $2, $3, $4)")
	_, err = stmt.Exec(user.ID, user.Usertype, "resgistration", user.Name)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer stmt.Close()
	c.Redirect("/regsuccess")
	return nil
}
func RegisterSuccessful(c *fiber.Ctx) error {
	return c.Render("registerationsuccesful", map[string]interface{}{
		"title": "Registeration Successful"})

}
func Logout(c *fiber.Ctx)error{
	cookie := c.Cookies("jwt")
	claims, err := middlewares.CookieGetData(cookie, c)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Not Logged In",
		})
	}

	if claims["name"] != "nologin" {
		c.ClearCookie()
		return c.Render("logout", map[string]interface{}{
			"title" : "Logout",
			"msg":  claims["name"],
			"logoutmsg":"Succesfully Logged out.",
		})
	} else {
		return c.Render("logout", map[string]interface{}{
			"title" : "Logout",
			"logoutmsg":"Not Logged In.",
		})
	}


}
