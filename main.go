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
	// Note: This is just an example, please use a secure secret key
	jwt := middlewares.NewAuthMiddleware(config.Secret)
	// Create a Login route
	app.Post("/login", handlers.Login)
	// Create a protected route
	app.Get("/protected", jwt, handlers.Protected)
	// Listen on port 3000
	app.Listen(":3000")
}
