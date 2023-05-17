package handlers

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/Praiseson6065/Golang_LibraryManagementSystem/config"
	"github.com/Praiseson6065/Golang_LibraryManagementSystem/middlewares"
	"github.com/Praiseson6065/Golang_LibraryManagementSystem/models"
	"github.com/Praiseson6065/Golang_LibraryManagementSystem/repository"
	"github.com/gofiber/fiber/v2"
	jtoken "github.com/golang-jwt/jwt/v4"
)

//home page
func HomePage(c *fiber.Ctx) error {
	return c.Render("static/home.html", map[string]interface{}{})
}

//profile page
func ProfilePage(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")
	token, err := jtoken.Parse(cookie, func(token *jtoken.Token) (interface{},error) {
		// Provide the secret key used for signing the token
		return []byte(config.Secret), nil
	})

	if err != nil {
		// Handle token parsing error
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid token",
		})
	}

	// Access the claims from the parsed token
	claims, ok := token.Claims.(jtoken.MapClaims)
	if !ok || !token.Valid {
		// Handle invalid token or invalid claims
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid token claims",
		})
	}

	return c.Render("static/profile.html",map[string]interface{}{
		"name" : claims["name"],
		"email" : claims["email"],
	})

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
		"ID":    user.ID,
		"email": user.Email,
		"name":  user.Name,
		"exp":   time.Now().Add(day * 1).Unix(),
	}

	// Create token
	token := jtoken.NewWithClaims(jtoken.SigningMethodHS256, claims)
	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte(config.Secret))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	c.Set("Authorization", "Bearer"+t)
	c.Cookie(&fiber.Cookie{
		Name:     "jwt",
		Value:    t,
		Expires:  time.Now().Add(24 * time.Hour),
		HTTPOnly: true,
		SameSite: "lax",
	})

	return c.Redirect("/profile")
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
	return c.Render("static/login.html", map[string]interface{}{
		"title": "Login Page"})
}

//register
func Register(c *fiber.Ctx) error {
	return c.Render("static/register.html", map[string]interface{}{
		"title": "Register Page"})
}
func RegisterPost(c *fiber.Ctx) error {
	db, err := sql.Open("postgres", "postgres://lib:lib@localhost:5432/lib")
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
	hashpassword,err:=middlewares.HashPassword(data.Password)
	if err != nil {
		return nil
	}
	_, err = stmt.Exec(data.Email, hashpassword , data.Name, "user")
	if err != nil {
		fmt.Println(err)
		return err
	}
	c.Redirect("/regsuccess")
	return nil
}
func RegisterSuccessful(c *fiber.Ctx) error {
	return c.Render("static/registerationsuccesful.html", map[string]interface{}{
		"title": "Registeration Successful"})

}
