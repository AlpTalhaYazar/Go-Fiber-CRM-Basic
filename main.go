package main

import (
	"github.com/gofiber/fiber"
	"github.com/talha-yazar/Go-Fiber-CRM-Basic/database"
)

func main() {
	app := fiber.New()
	app.Listen(3000)
	defer database.DBConn.Close()
}
