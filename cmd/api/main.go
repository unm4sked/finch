package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"log"
	"time"

	"github.com/unm4sked/finch/internal/api/routes"
	"github.com/unm4sked/finch/pkg/postgres"
)

func main() {
	fmt.Println("Application started...")

	database, err := postgres.New(postgres.MaxPoolSize(1), postgres.ConnTimeout(time.Second))
	if err != nil {
		log.Fatal(err)
	}
	defer database.Close()

	app := fiber.New()
	api := app.Group("/api").Group("/v1")
	routes.ConfigurationRouter(api)

	log.Fatal(app.Listen(":3000"))

}
