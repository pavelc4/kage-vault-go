package routes

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/pavelc4/kage-vault-go/config"
	"github.com/pavelc4/kage-vault-go/pkg/response"
)

var startTime = time.Now()

func Setup(app *fiber.App, cfg *config.Config) {
	app.Get("/", func(c *fiber.Ctx) error {
		return response.Success(c, fiber.Map{
			"service": cfg.Service,
			"version": cfg.Version,
		})
	})

	api := app.Group("/api")
	api.Get("/health", func(c *fiber.Ctx) error {
		return response.Success(c, fiber.Map{
			"status":  "healthy",
			"uptime":  time.Since(startTime).String(),
			"service": cfg.Service,
		})
	})
}
