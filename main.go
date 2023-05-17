package main

import (

	"github.com/Praiseson6065/Golang_LibraryManagementSystem/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	// Create a new Fiber instance
	app := fiber.New()
	
	app.Static("/", "./static")
	//routes
	
	routes.Setuproutes(app)	
	// Listen on port 3000
	app.Listen(":3000")

}
