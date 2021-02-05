package ranking

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"star-track.com/star-track-go-api/cors"
)

const rankingPath = "ranking"

func handleRanking(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		rankingList := getRankingList()
		j, err := json.Marshal(rankingList)
		if err != nil {
			log.Fatal(err)
		}
		_, err = w.Write(j)
		if err != nil {
			log.Fatal(err)
		}
	case http.MethodOptions:
		return
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

// SetupRoutes :
func SetupRoutes(apiBasePath string) {
	rankingHandler := http.HandlerFunc(handleRanking)
	http.Handle(fmt.Sprintf("%s/%s", apiBasePath, rankingPath), cors.Middleware(rankingHandler))
}
