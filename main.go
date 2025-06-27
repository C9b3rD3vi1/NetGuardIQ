package main

import (
	"fmt"
	fiber "github.com/gofiber/fiber/v2"

	"github.com/gofiber/template/html/v2"
	"github.com/c9b3rd3vi1/NetGuardIQ/handlers"
	"github.com/c9b3rd3vi1/NetGuardIQ/database"
	//"github.com/c9b3rd3vi1/NetGuardIQ/config"
	//"github.com/c9b3rd3vi1/NetGuardIQ/utils"
	//"github.com/c9b3rd3vi1/NetGuardIQ/models"

)


// main function initializes the Fiber application and sets up routes, templates, and static files.
func main() {

	// set template engine
	engine := html.New("./templates", ".html")


	app := fiber.New(fiber.Config{
		Views: engine,
	})
	// Enable Fiber's built-in logger
	//app.Use(fiber.Logger())

	// set database connection
	database.ConnectDB()

	// set static files directory
	app.Static("/", "./public")


	app.Get("/", handlers.Dashboard)
	app.Get("/campaigns/new", handlers.NewCampaignForm)
	app.Get("/campaigns/new", handlers.CreateCampaign)
	app.Get("/tracking/:id", handlers.TrackClick)
	app.Get("/fake_login", handlers.FakeLogin)
	app.Get("/login", handlers.UserLoginHandler)

	app.Listen(":3000")
	fmt.Println("Server is running on port 3000")
}