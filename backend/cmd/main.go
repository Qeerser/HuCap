package main

import (
	"HuCap/base/database"
	"HuCap/config"
	"HuCap/routes"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	
	if err := godotenv.Load() ; err != nil {
		log.Println("No .env file found")
	}
	cfg := config.LoadConfig()
	db,err := database.ConnectDB(cfg)
	
	if err != nil {
		log.Printf("Failed to connect to database: %v", err)
	}

	app := fiber.New()

	routes.SetupRoute(app,db)

	app.Listen(":" + cfg.ServerPort)
}
