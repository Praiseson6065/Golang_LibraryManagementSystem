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

	//api
	api := app.Group("/api")
	api.Post("/book", handlers.AddBooksPost)
	api.Get("/getbooks", handlers.GetBooks)
	api.Post("/searchbook", handlers.SearchBooks)
	api.Get("/book/:id", handlers.GetBook)
	api.Get("/bookc/:bc", handlers.GetBookByCode)
	api.Put("/updatebook/:id", handlers.UpdateBook)
	//like

	//cart
	api.Post("/cart/:userid/:bookid", handlers.AddtoCart)
	api.Get("/getusercart/:userid", handlers.GetUserCart)
	api.Delete("/cart/:userid/:bookid", handlers.RemoveFromCart)
	api.Post("/checkoutcart/:userid", handlers.CheckOutFromCart)
	api.Get("/cartbkchk/:userid/:bookid", handlers.ChkBookCart)
	//user
	api.Get("/issuedbooks/:userid", handlers.UserIssuedBooks)
	api.Get("/isbookIssued/:userid/:bookid", handlers.IsBookIssued)
	api.Post("/returnbook/:userid/:bookid", handlers.ReturnBook)
	api.Post("/like/:userid/:bookid", handlers.LikeBook)
	api.Get("/isliked/:userid/:bookid", handlers.IsLiked)
	api.Post("/reqbooks/",handlers.UserRequestedBooks)
	api.Get("/userreqbook/:userid",handlers.RequestedBooks)
	//admin
	admin := app.Group("/admin")
	admin.Get("/users", handlers.Userslist)
	admin.Post("/addadmin", handlers.AddAdmin)
	admin.Get("/reqbooks",handlers.ReqBook)
}
