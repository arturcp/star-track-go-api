package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

const basePath = "/api"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}

	SetupRoutes(basePath)
	fmt.Println("\nWebserver is up and running on port 5000, waiting for connections...")
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal(err)
	}
}
