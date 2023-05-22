package routes

import (
	"github.com/Praiseson6065/Golang_LibraryManagementSystem/config"
	"github.com/Praiseson6065/Golang_LibraryManagementSystem/handlers"
	"github.com/Praiseson6065/Golang_LibraryManagementSystem/middlewares"
	"github.com/gofiber/fiber/v2"
)

func Setuproutes(app *fiber.App) {

	jwt := middlewares.NewAuthMiddleware(config.Secret)

	app.Get("/", handlers.HomePage)
	app.Get("/profile", handlers.ProfilePage)
	app.Post("/login", handlers.Login)

	app.Get("/protected", jwt, handlers.Protected)

	app.Get("/login", handlers.Loginpage)

	app.Get("/register", handlers.Register)
	app.Post("/register", handlers.RegisterPost)
	app.Get("/regsuccess", handlers.RegisterSuccessful)
	app.Get("/logout", handlers.Logout)
	//admin
	app.Get("/admin", handlers.AdminPage)
}
