package dialog

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
)

type dialogList map[int][]Dialog
type dialogDictionary map[int]dialogList

var dialogsMap = struct {
	sync.RWMutex
	m map[int]dialogDictionary
}{m: make(map[int]dialogDictionary)}

func init() {
	fmt.Print("loading dialogs...")
	loadedMap, err := loadDialogsMap()
	dialogsMap.m = loadedMap
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d dialogs loaded...\n", len(dialogsMap.m))
}

func loadDialogsMap() (map[int]dialogDictionary, error) {
	newMap := map[int]dialogDictionary{}
	fmt.Println("")

	err := filepath.Walk("data/dialogs/",
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if filepath.Ext(path) == ".json" {
				fmt.Println("Reading file ", path)
				parts := strings.Split(path, "/")

				levelID, err := strconv.Atoi(strings.Replace(parts[2], "level-", "", 1))
				if err != nil {
					panic(err)
				}

				stageID, err := strconv.Atoi(strings.Replace(parts[3], "stage-", "", 1))
				if err != nil {
					panic(err)
				}

				if newMap[levelID] == nil {
					newMap[levelID] = dialogDictionary{}
				}

				if newMap[levelID][stageID] == nil {
					newMap[levelID][stageID] = dialogList{}
				}

				file, _ := ioutil.ReadFile(path)
				dialogs := newMap[levelID][stageID]
				err = json.Unmarshal([]byte(file), &dialogs)

				if err != nil {
					log.Fatal(err)
				}
			}
			return nil
		})
	if err != nil {
		log.Println(err)
	}

	return newMap, nil
}

func getDialog(levelID, stageID, dialogID int) []Dialog {
	dialogsMap.RLock()
	defer dialogsMap.RUnlock()

	if dialogs, ok := dialogsMap.m[levelID][stageID][dialogID]; ok {
		return dialogs
	}

	return []Dialog{}
}
