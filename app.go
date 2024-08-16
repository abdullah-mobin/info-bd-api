package main

import (
	"info-bd-api/database"
	"info-bd-api/routes"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	err := database.ConnectDB()
	if err != nil {
		log.Fatalf("failed to connect db", err)
	}

	if os.Getenv("APP_ENV") == "development" {
		database.MigrateSchema()
	}

	routes.SetUpRoutes(app)
	app.Listen(":3030")
}
