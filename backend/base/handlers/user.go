package handlers

import (
	"HuCap/base/schemas"
	"HuCap/base/services"

	"github.com/gofiber/fiber/v2"
)

type userHandler struct {
	userService services.UserService
}

func NewUserHandler(userService services.UserService) *userHandler {
	return &userHandler{
		userService: userService,
	}
}

func (h *userHandler) CreateUser(c *fiber.Ctx) error {
	var user schemas.User
	if err := c.BodyParser(&user) ; err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	
	if err := h.userService.CreateUser(&user) ; err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	return c.SendStatus(200)
}

func (h *userHandler) LoginUser(c *fiber.Ctx) error {
		var user schemas.User

		if err := c.BodyParser(&user); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
		}

		// Check if user exists
		if h.userService.LoginUser(&user) {
			return c.Status(fiber.StatusOK).JSON(fiber.Map{
				"message": "Login successful",
				"user":    user,
				"token":   "mock-jwt-token-12345", // Simulated JWT token
			})
		}

		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid email or password"})
}
