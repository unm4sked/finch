package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/unm4sked/finch/internal/api/handlers"
	"github.com/unm4sked/finch/internal/configuration"
)

func ConfigurationRouter(app fiber.Router, service configuration.Service) {
	app.Get("/configurations/:id", handlers.GetConfigurationById(service))
	app.Get("/configurations", handlers.GetConfigurations(service))
	app.Post("/configurations", handlers.CreateConfiguration(service))
	app.Patch("/configurations/:id", handlers.PatchConfiguration(service))
	app.Delete("/configuration/:id", handlers.DeleteConfiguration(service))
}
