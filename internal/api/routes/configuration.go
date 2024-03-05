package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/unm4sked/finch/internal/api/handlers"
)

func ConfigurationRouter(app fiber.Router) {
	app.Get("/configurations/:id", handlers.GetConfigurationById())
	app.Get("/configurations", handlers.GetConfigurations())
	app.Post("/configurations", handlers.CreateConfiguration())
	app.Patch("/configurations/:id", handlers.PatchConfiguration())
	app.Delete("/configuration/:id", handlers.DeleteConfiguration())
}
