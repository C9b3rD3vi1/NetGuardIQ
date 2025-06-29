package main

import (
	"fmt"
	fiber "github.com/gofiber/fiber/v2"

	"github.com/gofiber/template/html/v2"
	"github.com/c9b3rd3vi1/NetGuardIQ/handlers"
	"github.com/c9b3rd3vi1/NetGuardIQ/database"
	//"github.com/CloudyKit/jet/v6"
	//"github.com/c9b3rd3vi1/NetGuardIQ/config"
	//"github.com/c9b3rd3vi1/NetGuardIQ/utils"
	//"github.com/c9b3rd3vi1/NetGuardIQ/models"

)


// main function initializes the Fiber application and sets up routes, templates, and static files.
func main() {

	// set template engine
	engine := html.New("./templates", ".html")
	engine.Reload(true) // Disable template reloading in production mode
	

	app := fiber.New(fiber.Config{
		Views: engine,
		ViewsLayout: "/layouts/base",
	})


	// set database connection
	database.ConnectDB()

	// set static files directory
	app.Static("/public", "./public")

	// set routes
	app.Get("/", handlers.HomepageHandler)
	app.Get("/dashboard", handlers.Dashboard)
	app.Get("/campaigns/new", handlers.NewCampaignForm)
	app.Post("/campaigns/new", handlers.CreateCampaign)
	app.Get("/tracking/:id", handlers.TrackClick)
	app.Get("/fake_login", handlers.FakeLogin)
	app.Get("/login", handlers.UserLoginHandler)
	app.Get("/signup", handlers.UserRegisterHandler)
	app.Get("/pricing", handlers.PrinceHandler)
	app.Get("/products", handlers.ProductsHandler)

	
	fmt.Println("Server is running on port 3000")
	if err := app.Listen(":3000"); err != nil {
    fmt.Println("Error starting server:", err)
	}

}