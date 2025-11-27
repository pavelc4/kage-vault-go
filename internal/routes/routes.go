package routes

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/pavelc4/kage-vault-go/config"
	"github.com/pavelc4/kage-vault-go/internal/handlers"
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

	limiterConfig := limiter.New(limiter.Config{
		Max:        20,
		Expiration: 30 * time.Second,
		KeyGenerator: func(c *fiber.Ctx) string {
			return c.IP()
		},
		LimitReached: func(c *fiber.Ctx) error {
			return c.Status(fiber.StatusTooManyRequests).JSON(fiber.Map{
				"success": false,
				"message": "Too many requests, please try again later.",
			})
		},
	})
	api.Get("/password", limiterConfig, handlers.GeneratePasswordHandler)
}
