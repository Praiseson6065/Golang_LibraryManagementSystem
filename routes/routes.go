package routes

import (
	"github.com/Praiseson6065/Golang_LibraryManagementSystem/config"
	"github.com/Praiseson6065/Golang_LibraryManagementSystem/handlers"
	"github.com/Praiseson6065/Golang_LibraryManagementSystem/middlewares"
	"github.com/gofiber/fiber/v2"
)

func Setuproutes(app *fiber.App) {

	jwt := middlewares.NewAuthMiddleware(config.Secret)

	app.Get("/protected", jwt, handlers.Protected)
	//authorization
	app.Get("/login", handlers.Loginpage)
	app.Post("/login", handlers.Login)
	app.Post("/register", handlers.RegisterPost)
	app.Get("/regsuccess", handlers.RegisterSuccessful)
	app.Get("/logout", handlers.Logout)

	//testing

	//api
	api := app.Group("/api")

	api.Post("/book", handlers.AddBooksPost)
	api.Get("/getbooks", handlers.GetBooks)
	api.Post("/searchbook", handlers.SearchBooks)
	api.Get("/book/:id", handlers.GetBook)
	api.Get("/bookc/:bc", handlers.GetBookByCode)
	api.Get("/reviews/:bookid", handlers.BookReviewsByBookId)

	user := app.Group("/user")
	//userreview
	user.Post("/bookreview/:userid", middlewares.UserMiddleWare, handlers.BookReviewByUser)
	user.Delete("/delbookreview/:userid/:bookid", middlewares.UserMiddleWare, handlers.DeleteReviewByUser)
	user.Put("/updatereview/:userid", middlewares.UserMiddleWare, handlers.UpdateReviewByUser)
	//usercart
	user.Get("/profile", middlewares.UserMiddleWare, handlers.ProfilePage)
	user.Post("/cart/:userid/:bookid", middlewares.UserMiddleWare, handlers.AddtoCart)
	user.Get("/getusercart/:userid", middlewares.UserMiddleWare, handlers.GetUserCart)
	user.Delete("/cart/:userid/:bookid", middlewares.UserMiddleWare, handlers.RemoveFromCart)
	user.Post("/checkoutcart/:userid", middlewares.UserMiddleWare, handlers.CheckOutFromCart)
	user.Get("/cartbkchk/:userid/:bookid", middlewares.UserMiddleWare, handlers.ChkBookCart)
	//user
	user.Get("/issuedbooks/:userid", middlewares.UserMiddleWare, handlers.UserIssuedBooks)
	user.Get("/isbookIssued/:userid/:bookid", middlewares.UserMiddleWare, handlers.IsBookIssued)
	user.Delete("/issuebook/:userid/:bookid",middlewares.UserMiddleWare,handlers.UserRemoveIssueBook)
	user.Post("/returnbook/:userid/:bookid", middlewares.UserMiddleWare, handlers.ReturnBook)
	user.Post("/like/:userid/:bookid", middlewares.UserMiddleWare, handlers.LikeBook)
	user.Get("/isliked/:userid/:bookid", middlewares.UserMiddleWare, handlers.IsLiked)
	user.Post("/reqbooks/", handlers.UserRequestedBooks)
	user.Get("/userreqbook/:userid", middlewares.UserMiddleWare, handlers.RequestedBooks)
	user.Get("/approvbooks/:userid", middlewares.UserMiddleWare, handlers.UserApprovedBooks)
	user.Get("/userbookdetails/:userid/:bookid", middlewares.UserMiddleWare, handlers.UserBookDetails)

	//admin

	admin := app.Group("/admin")
	admin.Use(func(c *fiber.Ctx) error {
		cookie := c.Cookies("jwt")
		claims, _ := middlewares.CookieGetData(cookie, c)
		if claims["usertype"] != "admin" {
			return c.SendString("Un Authorized")
		}

		return c.Next()
	})
	admin.Put("/updatebook/:id", handlers.UpdateBook)
	admin.Get("/users", handlers.Userslist)
	admin.Post("/addadmin", handlers.AddAdmin)
	admin.Get("/reqbooks", handlers.ReqBook)
	admin.Post("/approvbooks", handlers.ApprovBookByAdmin)
	admin.Get("/approvalbookslist", handlers.GetApprovalBooks)
}
