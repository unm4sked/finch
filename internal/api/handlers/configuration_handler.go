package handlers

import (
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/unm4sked/finch/internal/api/responses"
	"github.com/unm4sked/finch/internal/configurations"
)

func GetConfigurationById(service configurations.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")
		configuration, err := service.GetConfiguration(id)
		if err != nil {
			return c.Status(http.StatusBadRequest).JSON(responses.NewFailureResponseBody(err))
		}
		return c.Status(http.StatusBadGateway).JSON(responses.NewSuccessResponseBody(configuration))
	}
}

func GetConfigurations(service configurations.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		configs, err := service.GetConfigurations()
		if err != nil {
			return c.Status(http.StatusBadRequest).JSON(responses.NewFailureResponseBody(err))
		}
		return c.Status(http.StatusBadGateway).JSON(responses.NewSuccessResponseBody(configs))
	}
}

func DeleteConfiguration(service configurations.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")
		err := service.RemoveConfiguration(id)
		if err != nil {
			return c.Status(http.StatusBadRequest).JSON(responses.NewFailureResponseBody(err))
		}
		return c.Status(http.StatusBadGateway).JSON(responses.NewSuccessResponseBody(fiber.Map{}))
	}
}

func PatchConfiguration(service configurations.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		updateConfigurationPayload := struct {
			Description string `json:"description"`
		}{}

		if err := c.BodyParser(&updateConfigurationPayload); err != nil {
			fmt.Println("ja jebie xd")
			return c.Status(http.StatusBadRequest).JSON(responses.NewFailureResponseBody(err))
		}
		id := c.Params("id")

		err := service.UpdateConfigurationDescription(id, updateConfigurationPayload.Description)
		if err != nil {
			return c.Status(http.StatusBadRequest).JSON(responses.NewFailureResponseBody(err))
		}
		return c.Status(http.StatusOK).JSON(responses.NewSuccessResponseBody(fiber.Map{}))
	}
}

func CreateConfiguration(service configurations.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		createConfigurationPayload := struct {
			Description string `json:"description"`
		}{}
		if err := c.BodyParser(&createConfigurationPayload); err != nil {
			return c.Status(http.StatusBadRequest).JSON(responses.NewFailureResponseBody(err))
		}
		id, err := service.Create(createConfigurationPayload.Description)
		if err != nil {
			return c.Status(http.StatusBadRequest).JSON(responses.NewFailureResponseBody(err))
		}
		return c.Status(http.StatusCreated).JSON(responses.NewSuccessResponseBody(fiber.Map{
			"id": id,
		}))
	}
}
