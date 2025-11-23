package main

import (
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
)

var (
	startTime = time.Now()
	version   = "1.0.0"
	service   = "KageVault API"
)

func main() {
	app := fiber.New(fiber.Config{
		AppName: service + " v" + version,
	})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"success":   true,
			"service":   service,
			"version":   version,
			"message":   "API is running",
			"timestamp": time.Now().Format(time.RFC3339),
		})
	})

	app.Get("/api/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"success":   true,
			"service":   service,
			"version":   version,
			"status":    "healthy",
			"uptime":    time.Since(startTime).String(),
			"timestamp": time.Now().Format(time.RFC3339),
		})
	})

	log.Println("🚀", service, "running at http://localhost:3000")
	log.Fatal(app.Listen(":3000"))
}
