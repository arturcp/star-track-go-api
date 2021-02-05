package main

import (
	"fmt"
	"net/http"

	"star-track.com/star-track-go-api/character"
	"star-track.com/star-track-go-api/cors"
	"star-track.com/star-track-go-api/ranking"
)

const charactersPath = "characters"
const rankingPath = "ranking"

// SetupRoutes configures the routes for the API.
func SetupRoutes(apiBasePath string) {
	rankingHandler := http.HandlerFunc(ranking.HandleRanking)
	http.Handle(fmt.Sprintf("%s/%s", apiBasePath, rankingPath), cors.Middleware(rankingHandler))

	charactersHandler := http.HandlerFunc(character.HandleCharacters)
	characterHandler := http.HandlerFunc(character.HandleCharacter)
	http.Handle(fmt.Sprintf("%s/%s", apiBasePath, charactersPath), cors.Middleware(charactersHandler))
	http.Handle(fmt.Sprintf("%s/%s/", apiBasePath, charactersPath), cors.Middleware(characterHandler))
}
