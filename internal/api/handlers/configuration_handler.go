package handlers

import "github.com/gofiber/fiber/v2"

func GetConfigurationById() fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.JSON([]int{})
	}
}
