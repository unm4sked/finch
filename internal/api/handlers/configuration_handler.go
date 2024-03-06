package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/unm4sked/finch/internal/configuration"
)

func GetConfigurationById(service configuration.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")
		configuration, err := service.GetConfiguration(id)
		if err != nil {
			return c.Status(400).JSON(fiber.Map{
				"error": err,
			})
		}
		return c.JSON(configuration)
	}
}

func GetConfigurations(service configuration.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		configs, err := service.GetConfigurations()
		if err != nil {
			return c.Status(400).JSON(fiber.Map{
				"error": err,
			})
		}
		return c.JSON(configs)
	}
}

func DeleteConfiguration(service configuration.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")
		err := service.RemoveConfiguration(id)
		if err != nil {
			return c.Status(400).JSON(fiber.Map{
				"error": err,
			})
		}
		return c.JSON(fiber.Map{
			"status": "removed",
		})
	}
}

func PatchConfiguration(service configuration.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		type UpdateConfiguration struct {
			Description string `json:"description"`
		}
		updateConfig := new(UpdateConfiguration)

		if err := c.BodyParser(updateConfig); err != nil {
			return c.Status(400).JSON(fiber.Map{
				"error": err,
			})
		}
		id := c.Params("id")

		err := service.UpdateConfigurationDescription(id, updateConfig.Description)
		if err != nil {
			return c.Status(400).JSON(fiber.Map{
				"error": err,
			})
		}
		return c.JSON(fiber.Map{
			"status": "removed",
		})
	}
}

func CreateConfiguration(service configuration.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		service.Create()

		return c.JSON([]string{})
	}
}
