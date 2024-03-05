package main

import (
	"fmt"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"

	"github.com/unm4sked/finch/internal/api/routes"
	"github.com/unm4sked/finch/internal/configuration"
	"github.com/unm4sked/finch/pkg/postgres"
)

func main() {
	fmt.Println("Application started...")

	database, err := postgres.New(postgres.MaxPoolSize(1), postgres.ConnTimeout(time.Second))
	if err != nil {
		log.Fatal(err)
	}
	defer database.Close()

	configurationService := configuration.NewService(configuration.NewPostgresRepository(*database))

	app := fiber.New()
	api := app.Group("/api").Group("/v1")

	routes.ConfigurationRouter(api, configurationService)

	log.Fatal(app.Listen(":3000"))

}
