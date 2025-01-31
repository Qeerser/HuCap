package routes

import (
	"HuCap/base/handlers"
	"HuCap/base/repositories"
	"HuCap/base/services"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// ฟังก์ชันนี้จะตั้งค่า API Routes
func SetupRoute(app *fiber.App, db *gorm.DB) {

	repository := repositories.NewRepository(db)
	service := services.NewService(repository)
	handler := handlers.NewHandler(service)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})
	
	api := app.Group("/api")

	api.Post("user" , handler.User().CreateUser)
}