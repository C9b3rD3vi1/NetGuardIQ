package handlers


import (
	//"github.com/c9b3rd3vi1/NetGuardIQ/database"
	//"github.com/c9b3rd3vi1/NetGuardIQ/models"
	fiber "github.com/gofiber/fiber/v2"
)


// User login Function
func UserLoginHandler(c *fiber.Ctx) error {
	// Render the login page
	return c.Render("login", fiber.Map{
		
	})
}

// User registration Function
func UserRegisterHandler(c *fiber.Ctx) error {
	// Render the registration page
	return c.Render("signup", fiber.Map{
		
	})
}