package handlers

import (
	"github.com/c9b3rd3vi1/NetGuardIQ/database"
	"github.com/c9b3rd3vi1/NetGuardIQ/models"
	fiber "github.com/gofiber/fiber/v2"
)


// TrackClick handles the tracking of clicks on phishing links
func TrackClick(c *fiber.Ctx) error {
	id := c.Params("id")
	var target models.Target

	// Find the target by ID
	if err := database.DB.First(&target, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).SendString("Target not found")
	}

	// Update the Clicked field to true
	target.Clicked = true
	if err := database.DB.Save(&target).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Error updating target")
	}

	// Redirect to the phishing page or a thank you page
	return c.Redirect("/fake_login")
}

// FakeLogin simulates a phishing login page
func FakeLogin(c *fiber.Ctx) error {
	// Render the fake login page
	return c.Render("fake_login", fiber.Map{
		//"Title": "Fake Login",
	})
}