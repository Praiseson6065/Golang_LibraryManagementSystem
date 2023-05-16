package main

import (
	"github.com/Praiseson6065/Golang_LibraryManagementSystem/config"
	"github.com/Praiseson6065/Golang_LibraryManagementSystem/handlers"
	"github.com/Praiseson6065/Golang_LibraryManagementSystem/middlewares"
	"github.com/gofiber/fiber/v2"
)

func main() {
	// Create a new Fiber instance
	app := fiber.New()
	// Create a new JWT middleware
	jwt := middlewares.NewAuthMiddleware(config.Secret)
	app.Static("/", "./static")
	// Create a Login route
	app.Post("/login", handlers.Login)
	// Create a protected route
	app.Get("/protected", jwt, handlers.Protected)

	app.Get("/loginpage", handlers.Loginpage)
	app.Post("/loginpage", handlers.LoginpagePost)
	
	//registerpage
	app.Get("/register", handlers.Register)
	app.Post("/register",handlers.RegisterPost)
	app.Get("/regsuccess",handlers.RegisterSuccessful)
	// Listen on port 3000
	app.Listen(":3000")	

}
