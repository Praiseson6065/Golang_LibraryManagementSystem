package main

import (
	"github.com/Praiseson6065/Golang_LibraryManagementSystem/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
	
)

func main() {
	engine := html.New("./static", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})
	app.Static("/", "./static", fiber.Static{Index: "home.html"})

	routes.Setuproutes(app)

	// data, _ := json.MarshalIndent(app.Stack(), "", "  ")
	// fmt.Println(string(data))
	app.Listen(":3000")
	
}
