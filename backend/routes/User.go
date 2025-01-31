package routes

import (
	"HuCap/base/handlers"
	"HuCap/base/repositories"
	"HuCap/base/services"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"gorm.io/gorm"
)

// ฟังก์ชันนี้จะตั้งค่า API Routes
func SetupRoute(app *fiber.App, db *gorm.DB) {

	repository := repositories.NewRepository(db)
	service := services.NewService(repository)
	handler := handlers.NewHandler(service)

	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:3000", // Allow requests from Next.js frontend
		AllowMethods: "GET,POST,PUT,DELETE,OPTIONS",
		AllowHeaders: "Content-Type, Authorization",
	}))

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})
	
	api := app.Group("/api")

	api.Post("register" , handler.User().CreateUser)
	api.Post("login" , handler.User().LoginUser)
}