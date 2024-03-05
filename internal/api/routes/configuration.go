package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/unm4sked/finch/internal/api/handlers"
)

func ConfigurationRouter(app fiber.Router) {
	app.Get("/configurations/:id", handlers.GetConfigurationById())
}
