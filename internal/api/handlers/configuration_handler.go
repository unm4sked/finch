package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/unm4sked/finch/internal/configuration"
)

func GetConfigurationById(service configuration.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.JSON([]string{})
	}
}

func GetConfigurations(service configuration.Service) fiber.Handler {
	return func(c *fiber.Ctx) error { return c.JSON([]string{}) }
}

func DeleteConfiguration(service configuration.Service) fiber.Handler {
	return func(c *fiber.Ctx) error { return c.JSON([]string{}) }
}

func PatchConfiguration(service configuration.Service) fiber.Handler {
	return func(c *fiber.Ctx) error { return c.JSON([]string{}) }
}

func CreateConfiguration(service configuration.Service) fiber.Handler {
	return func(c *fiber.Ctx) error { return c.JSON([]string{}) }
}
