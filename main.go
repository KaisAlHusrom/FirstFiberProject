package main

import (
	"log"
	"os"

	Config "github.com/KaisAlHusrom/FirstFiberProject/Config"
	Routes "github.com/KaisAlHusrom/FirstFiberProject/Routes"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	//connect with database
	Config.Connect()

	//download env file
	godotenv.Load(".env")

	//new fiber app
	app := fiber.New()

	portString := os.Getenv("PORT")
	if portString == "" {
		log.Fatal("PORT is not found in the environment")
	}

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{
			"success": true,
			"message": "First api created",
		})
	})

	log.Fatal(app.Listen(":" + portString))
}
