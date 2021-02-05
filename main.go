package main

import (
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
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal(err)
	}
}
