package main

import (
	"log"
	"net/http"
	"os"

	"azure-service/routes"
)

func main() {
	port := os.Getenv("HTTP_PLATFORM_PORT")

	if port == "" {
		port = "8080"
	}

	err := http.ListenAndServe(":"+port, routes.Init())

	if err != nil {
		log.Fatal(err)
	}
}
