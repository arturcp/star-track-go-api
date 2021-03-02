package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	"star-track.com/star-track-go-api/character"
	"star-track.com/star-track-go-api/cors"
	"star-track.com/star-track-go-api/dialog"
	"star-track.com/star-track-go-api/level"
	"star-track.com/star-track-go-api/ranking"
)

const charactersPath = "characters"
const rankingPath = "ranking"

// SetupRoutes configures the routes for the API.
func SetupRoutes(apiBasePath string) {
	router := mux.NewRouter()
	router.Use(cors.Middleware)
	apiRouter := router.PathPrefix("/api").
		Methods("GET", "POST", "PUT", "DELETE", "OPTIONS").
		Subrouter()

	apiRouter.HandleFunc("/characters/{id}", character.HandleCharacter)
	apiRouter.HandleFunc("/characters", character.HandleCharacters)

	apiRouter.HandleFunc("/ranking", ranking.HandleRanking)

	apiRouter.HandleFunc("/dialogs/{levelID}/{stageID}/{id}", dialog.HandleDialogs)

	apiRouter.HandleFunc("/levels", level.HandleLevels)

	http.Handle("/", router)

	fmt.Print("Available routes:\n\n")
	router.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
		tpl, _ := route.GetPathTemplate()
		met, _ := route.GetMethods()
		fmt.Println(tpl, "", met, "")
		return nil
	})
}
