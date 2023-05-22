package main

import (

	"github.com/Praiseson6065/Golang_LibraryManagementSystem/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
)

func main() {
	// Create a new Fiber instance
	
	engine := html.New("./static", ".html")
	app := fiber.New(fiber.Config{

		Views: engine, // new config

	})
	app.Static("/", "./static")
	routes.Setuproutes(app)	
	// Listen on port 3000
	app.Listen(":3000")

}	
