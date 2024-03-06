package handlers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/unm4sked/finch/internal/configuration"
)

func GetConfigurationById(service configuration.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.JSON([]string{})
	}
}

func GetConfigurations(service configuration.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		configs, err := service.GetConfigurations()
		if err != nil {
			return c.SendStatus(500)
		}
		return c.JSON(configs)
	}
}

func DeleteConfiguration(service configuration.Service) fiber.Handler {
	return func(c *fiber.Ctx) error { return c.JSON([]string{}) }
}

func PatchConfiguration(service configuration.Service) fiber.Handler {
	return func(c *fiber.Ctx) error { return c.JSON([]string{}) }
}

func CreateConfiguration(service configuration.Service) fiber.Handler {
	fmt.Println("CreateConfiguration")
	return func(c *fiber.Ctx) error {
		service.Create()

		return c.JSON([]string{})
	}
}
