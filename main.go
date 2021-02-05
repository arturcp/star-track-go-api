package main

import (
	"log"
	"net/http"
	"os"

	"star-track.com/star-track-go-api/character"
)

const basePath = "/api"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}

	character.SetupRoutes(basePath)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal(err)
	}
}
