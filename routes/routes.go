package routes

import (
	"github.com/Praiseson6065/Golang_LibraryManagementSystem/config"
	"github.com/Praiseson6065/Golang_LibraryManagementSystem/handlers"
	"github.com/Praiseson6065/Golang_LibraryManagementSystem/middlewares"
	"github.com/gofiber/fiber/v2"
)

func Setuproutes(app *fiber.App) {

	jwt := middlewares.NewAuthMiddleware(config.Secret)

	//user
	app.Get("/profile", handlers.ProfilePage)

	app.Get("/protected", jwt, handlers.Protected)
	//authorization
	app.Get("/login", handlers.Loginpage)
	app.Post("/login", handlers.Login)
	app.Post("/register", handlers.RegisterPost)
	app.Get("/regsuccess", handlers.RegisterSuccessful)
	app.Get("/logout", handlers.Logout)
	//admin
	app.Get("/admin", handlers.AdminPage)
	//api
	api := app.Group("/api")
	api.Post("/book", handlers.AddBooksPost)
	api.Get("/getbooks", handlers.GetBooks)
	api.Post("/searchbook", handlers.SearchBooks)
	api.Get("/book/:id", handlers.GetBook)
	api.Put("/updatebook/:id", handlers.UpdateBook)
}
