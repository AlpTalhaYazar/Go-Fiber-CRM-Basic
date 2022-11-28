package main

import (
	"github.com/gofiber/fiber"
	"github.com/talha-yazar/Go-Fiber-CRM-Basic/database"
	"github.com/talha-yazar/Go-Fiber-CRM-Basic/lead"
)

// setupRoutes is a function that sets up all the routes for the application
func setupRoutes(app *fiber.App) {
	app.Get("/api/v1/lead", lead.GetLeads)
	app.Get("/api/v1/lead/:id", lead.GetLead)
	app.Post("/api/v1/lead", lead.NewLead)
	app.Put("/api/v1/lead/:id", lead.UpdateLead)
	app.Delete("/api/v1/lead/:id", lead.DeleteLead)
}

func main() {
	app := fiber.New()
	setupRoutes(app)
	app.Listen(3000)
	defer database.DBConn.Close()
}
