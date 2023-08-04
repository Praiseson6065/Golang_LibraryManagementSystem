package main

import (
	"fmt"

	"github.com/Praiseson6065/Golang_LibraryManagementSystem/config"
	"github.com/Praiseson6065/Golang_LibraryManagementSystem/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {

	config.InitEnvConfigs()
	fmt.Println(config.EnvConfigs)
	app := fiber.New()
	app.Static("/", "./static", fiber.Static{Index: "home.html"})

	routes.Setuproutes(app)

	// data, _ := json.MarshalIndent(app.Stack(), "", "  ")
	// fmt.Println(string(data))

	app.Listen(":3000")

}
