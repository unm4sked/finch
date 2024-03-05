package handlers

import "github.com/gofiber/fiber/v2"

func GetConfigurationById() fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.JSON([]string{})
	}
}

func GetConfigurations() fiber.Handler {
	return func(c *fiber.Ctx) error { return c.JSON([]string{}) }
}

func DeleteConfiguration() fiber.Handler {
	return func(c *fiber.Ctx) error { return c.JSON([]string{}) }
}

func PatchConfiguration() fiber.Handler {
	return func(c *fiber.Ctx) error { return c.JSON([]string{}) }
}

func CreateConfiguration() fiber.Handler {
	return func(c *fiber.Ctx) error { return c.JSON([]string{}) }
}
