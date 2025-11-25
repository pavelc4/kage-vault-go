package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/pavelc4/kage-vault-go/internal/models"
	"github.com/pavelc4/kage-vault-go/internal/services"
)

func GeneratePasswordHandler(c *fiber.Ctx) error {
	length := c.QueryInt("length", 12)
	useDigits := c.QueryBool("digits", true)
	useSymbols := c.QueryBool("symbols", true)
	useUpper := c.QueryBool("uppercase", true)

	config := services.PasswordConfig{
		Length:     length,
		UseLetters: true,
		UseDigits:  useDigits,
		UseSymbols: useSymbols,
		UseUpper:   useUpper,
	}

	password, err := services.GeneratePassword(config)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.ApiResponse{
			Success: false,
			Message: "Failed to generate password",
		})
	}

	return c.JSON(models.ApiResponse{
		Success: true,
		Message: "Password generated successfully",
		Data: fiber.Map{
			"password": password,
			"length":   length,
		},
	})
}
