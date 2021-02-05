package main

import (
	"log"
	"net/http"

	"star-track.com/star-track-go-api/character"
)

const basePath = "/api"

func main() {
	character.SetupRoutes(basePath)
	err := http.ListenAndServe(":5000", nil)
	if err != nil {
		log.Fatal(err)
	}
}
