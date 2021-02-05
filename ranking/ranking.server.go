package ranking

import (
	"encoding/json"
	"log"
	"net/http"
)

func HandleRanking(w http.ResponseWriter, r *http.Request) {
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
