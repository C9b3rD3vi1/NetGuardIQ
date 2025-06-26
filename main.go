package main

import (
	"fmt"
	"os"
	fiber "github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/template/html"

	"github.com/c9b3rd3vi1/Golang/NetGuardIQ/database"
)

func main() {

	// set template engine
	engine := html.New("./templates", ".html")


	app := fiber.New(fiber.Config{
		Views: engine,
	})

	// set database connection
	database.ConnectDB()

	// set static files directory
	app.Static("/", "./public")


	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})
	app.Listen(":3000")
	fmt.Println("Server is running on port 3000")
}