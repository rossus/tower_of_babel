package converter

import (
	"io/ioutil"
	"github.com/rossus/tower_of_babel/common/types"
	"strings"
	"encoding/json"
	"github.com/rossus/tower_of_babel/session"
	"os"
)

func checkFilename(name string) bool {
	if strings.HasSuffix(name, ".json") {
		return true
	}
	return false
}

func RenewSessionInfoList() error {
	var sessionList = make(map[string]types.SessionInfo)

	filesInfo, err := ioutil.ReadDir("./saves")
	if err != nil {
		return err
	}
	for _, fileInfo := range filesInfo {
		if checkFilename(fileInfo.Name()) {
			var loadedSession types.LoadedSession
			var sessionInfo types.SessionInfo

			savedSessionJSON, err := ioutil.ReadFile("./saves/" + fileInfo.Name())
			if err != nil {
				return err
			}

			err = json.Unmarshal(savedSessionJSON, &loadedSession)
			if err == nil && loadedSession.Name != "" && loadedSession.Year > 0 && loadedSession.Version != "" {
				sessionInfo.Year = loadedSession.Year - 1
				sessionInfo.Name = loadedSession.Name
				sessionInfo.Version = loadedSession.Version

				sessionList[sessionInfo.Name] = sessionInfo
			} else {
				errR := os.Rename("./saves/" + fileInfo.Name(), "./saves/old/" + fileInfo.Name())

				if errR != nil {
					return errR
				}
			}
		}
	}

	session.SetSessionList(sessionList)

	return nil
}
