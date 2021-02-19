package dialog

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sync"
)

var dialogsMap = struct {
	sync.RWMutex
	m map[int][]Dialog
}{m: make(map[int][]Dialog)}

func init() {
	fmt.Print("loading dialogs...")
	loadedMap, err := loadDialogsMap()
	dialogsMap.m = loadedMap
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d dialogs loaded...\n", len(dialogsMap.m))
}

func loadDialogsMap() (map[int][]Dialog, error) {
	fileName := "data/dialogs.json"
	_, err := os.Stat(fileName)
	if os.IsNotExist(err) {
		return nil, fmt.Errorf("file [%s] does not exist", fileName)
	}

	file, _ := ioutil.ReadFile(fileName)
	newMap := map[int][]Dialog{}
	err = json.Unmarshal([]byte(file), &newMap)

	if err != nil {
		log.Fatal(err)
	}

	return newMap, nil
}

func getDialog(dialogID int) []Dialog {
	dialogsMap.RLock()
	defer dialogsMap.RUnlock()

	if dialogs, ok := dialogsMap.m[dialogID]; ok {
		return dialogs
	}

	return []Dialog{}
}
