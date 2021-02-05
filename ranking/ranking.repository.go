package ranking

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sort"
	"sync"
)

var rankingMap = struct {
	sync.RWMutex
	m map[int]Ranking
}{m: make(map[int]Ranking)}

func init() {
	fmt.Println("loading ranking...")
	charMap, err := loadRankingMap()
	rankingMap.m = charMap
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("%d ranking loaded...\n", len(rankingMap.m))
}

func loadRankingMap() (map[int]Ranking, error) {
	fileName := "data/ranking.json"
	_, err := os.Stat(fileName)
	if os.IsNotExist(err) {
		return nil, fmt.Errorf("file [%s] does not exist", fileName)
	}

	file, _ := ioutil.ReadFile(fileName)
	rankingList := make([]Ranking, 0)
	err = json.Unmarshal([]byte(file), &rankingList)

	if err != nil {
		log.Fatal(err)
	}

	charMap := make(map[int]Ranking)
	for i := 0; i < len(rankingList); i++ {
		charMap[rankingList[i].ID] = rankingList[i]
	}

	return charMap, nil
}

func getRankingList() []Ranking {
	rankingMap.RLock()

	keys := make([]int, 0, len(rankingMap.m))
	for k := range rankingMap.m {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	ranking := make([]Ranking, 0, len(rankingMap.m))
	for _, key := range keys {
		ranking = append(ranking, rankingMap.m[key])
	}
	rankingMap.RUnlock()
	return ranking
}
