package level

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sort"
	"sync"
)

var levelsMap = struct {
	sync.RWMutex
	m map[int]Level
}{m: make(map[int]Level)}

func init() {
	fmt.Print("loading levels and stages...")
	loadedMap, err := loadLevelsMap()
	levelsMap.m = loadedMap
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d levels and stages loaded...\n", len(levelsMap.m))
}

func loadLevelsMap() (map[int]Level, error) {
	fileName := "data/levels_and_stages.json"
	_, err := os.Stat(fileName)
	if os.IsNotExist(err) {
		return nil, fmt.Errorf("file [%s] does not exist", fileName)
	}

	file, _ := ioutil.ReadFile(fileName)
	levelsList := make([]Level, 0)
	err = json.Unmarshal([]byte(file), &levelsList)

	if err != nil {
		log.Fatal(err)
	}

	newMap := make(map[int]Level)
	for i := 0; i < len(levelsList); i++ {
		newMap[levelsList[i].ID] = levelsList[i]
	}

	return newMap, nil
}

func getLevelsList() []Level {
	levelsMap.RLock()

	keys := make([]int, 0, len(levelsMap.m))
	for k := range levelsMap.m {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	levels := []Level{}
	for _, key := range keys {
		levels = append(levels, levelsMap.m[key])
	}
	levelsMap.RUnlock()
	return levels
}
