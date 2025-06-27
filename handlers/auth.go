package handlers


import (
	//"github.com/c9b3rd3vi1/NetGuardIQ/database"
	//"github.com/c9b3rd3vi1/NetGuardIQ/models"
	fiber "github.com/gofiber/fiber/v2"
)



func UserLoginHandler(c *fiber.Ctx) error {
	// Render the login page
	return c.Render("base", fiber.Map{
		"Title": "Login",
	})
}