package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/pavelc4/kage-vault-go/config"
	"github.com/pavelc4/kage-vault-go/internal/routes"
	"github.com/pavelc4/kage-vault-go/pkg/response"
)

func main() {
	cfg := config.Load()
	app := fiber.New(fiber.Config{
		AppName: cfg.Service,
	})

	routes.Setup(app, cfg)

	app.Use(func(c *fiber.Ctx) error {
		return response.Error(c, 404, "Route not found")
	})

	log.Printf("%s running on port %s", cfg.Service, cfg.Port)
	log.Fatal(app.Listen(":" + cfg.Port))
}
