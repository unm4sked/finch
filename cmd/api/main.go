package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/unm4sked/finch/pkg/postgres"
)

func main() {
	fmt.Println("Application started...")

	database, err := postgres.New()
	if err != nil {
		log.Fatal(err)
	}
	defer database.Close()

	app := fiber.New()
	app.Group("/api").Group("/v1")

	log.Fatal(app.Listen(":3000"))

}
