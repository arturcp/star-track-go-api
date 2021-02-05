package character

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sort"
	"sync"
)

var characterMap = struct {
	sync.RWMutex
	m map[int]Character
}{m: make(map[int]Character)}

func init() {
	fmt.Println("loading characters...")
	charMap, err := loadCharacterMap()
	characterMap.m = charMap
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("%d characters loaded...\n", len(characterMap.m))
}

func loadCharacterMap() (map[int]Character, error) {
	fileName := "data/characters.json"
	_, err := os.Stat(fileName)
	if os.IsNotExist(err) {
		return nil, fmt.Errorf("file [%s] does not exist", fileName)
	}

	file, _ := ioutil.ReadFile(fileName)
	charactersList := make([]Character, 0)
	err = json.Unmarshal([]byte(file), &charactersList)

	if err != nil {
		log.Fatal(err)
	}

	charMap := make(map[int]Character)
	for i := 0; i < len(charactersList); i++ {
		charMap[charactersList[i].ID] = charactersList[i]
	}

	return charMap, nil
}

func getCharacter(characterId int) *Character {
	characterMap.RLock()
	defer characterMap.RUnlock()

	if character, ok := characterMap.m[characterId]; ok {
		return &character
	}

	return nil
}

// func removeCharacter(characterId int) {
// 	characterMap.Lock()
// 	defer characterMap.Unlock()
// 	delete(characterMap.m, characterId)
// }

func getCharactersList() []Character {
	characterMap.RLock()

	keys := make([]int, 0, len(characterMap.m))
	for k := range characterMap.m {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	characters := make([]Character, 0, len(characterMap.m))
	for _, key := range keys {
		characters = append(characters, characterMap.m[key])
	}
	characterMap.RUnlock()
	return characters
}

func getCharactersIds() []int {
	characterMap.RLock()
	characterIds := []int{}
	for key := range characterMap.m {
		characterIds = append(characterIds, key)
	}
	characterMap.RUnlock()
	sort.Ints(characterIds)
	return characterIds
}

func getNextCharacterID() int {
	characterIds := getCharactersIds()
	return characterIds[len(characterIds)-1] + 1
}

func addOrUpdateCharacter(character Character) (int, error) {
	addOrUpdateID := -1
	if character.ID > 0 {
		oldCharacter := getCharacter(character.ID)
		if oldCharacter == nil {
			return 0, fmt.Errorf("character id [%d] doesn't exist", character.ID)
		}
		addOrUpdateID = character.ID
	} else {
		addOrUpdateID = getNextCharacterID()
		character.ID = addOrUpdateID
	}
	characterMap.Lock()
	characterMap.m[addOrUpdateID] = character
	characterMap.Unlock()
	return addOrUpdateID, nil
}
