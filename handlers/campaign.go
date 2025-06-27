package handlers

import (
	"github.com/c9b3rd3vi1/NetGuardIQ/database"
	"github.com/c9b3rd3vi1/NetGuardIQ/models"
	fiber "github.com/gofiber/fiber/v2"
)

// Homepage Handler
func HomepageHandler(c *fiber.Ctx) error {
	return c.Render("home", fiber.Map{
		"Title": "NetGuardIQ",
	})
}

func PrinceHandler(c *fiber.Ctx) error {
	return c.Render("pricing", fiber.Map{
		"Title": "Prince Handler",
	})
}


func ProductsHandler(c *fiber.Ctx) error {
	return c.Render("products", fiber.Map{
		"Title": "Products Handler",
	})
}


// Create Campaign Dashboard

func Dashboard(c *fiber.Ctx) error {
	var campaigns []models.Campaign
	if err := database.DB.Preload("Target").Find(&campaigns).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Error fetching campaigns")
	}
	return c.Render("dashboard", fiber.Map{
		"Title":     "Dashboard",
		"Campaigns": campaigns,
	})
}

// Create New Campaign Form
func NewCampaignForm(c *fiber.Ctx) error {
	return c.Render("base", fiber.Map{
		"Title": "New Campaign",
	})
}

// Create Campaign in Database
func CreateCampaign(c *fiber.Ctx) error {
	name := c.FormValue("name")
	description := c.FormValue("description")
	campaign := models.Campaign{
		Name:        name,
		Description: description,
	}
	// Save the campaign to the database
	if err := database.DB.Create(&campaign).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Error creating campaign")
	}
	return c.Redirect("/campaigns/new")
}