package level

import (
	"encoding/json"
	"log"
	"net/http"
)

func HandleLevels(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		levelsList := getLevelsList()
		j, err := json.Marshal(levelsList)
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
