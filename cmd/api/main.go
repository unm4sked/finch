package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/unm4sked/finch/pkg/postgres"
)

func main() {
	fmt.Println("Application started...")

	database, err := postgres.New()
	if err != nil {
		log.Fatal(err)
	}

	mux := http.NewServeMux()
	// mux.HandleFunc()

	http.ListenAndServe(":8080", mux)

	defer database.Close()
}
