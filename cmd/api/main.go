package main

import (
	"fmt"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
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
	app.Group("/api").Group("/v1")

	log.Fatal(app.Listen(":3000"))

}
