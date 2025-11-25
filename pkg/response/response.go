package response

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/pavelc4/kage-vault-go/internal/models"
)

func Success(c *fiber.Ctx, data interface{}) error {
	return c.JSON(models.ApiResponse{
		Success:   true,
		Data:      data,
		Timestamp: time.Now().Format(time.RFC3339),
	})
}

func Error(c *fiber.Ctx, statusCode int, message string) error {
	return c.Status(statusCode).JSON(models.ApiResponse{
		Success: false,
		Error: &models.ErrorDetail{
			Code:    statusCode,
			Message: message,
		},
		Timestamp: time.Now().Format(time.RFC3339),
	})
}
