package session

import (
	"github.com/rossus/tower_of_babel/common/types"
)

var blank = types.BlankSession{}
var active = types.ActiveSession{blank}
var sessionList = make(map[string]*types.Session)

func SetActiveSession(session *types.Session) {
	active.Sessions = session
}

func UnloadSession() {
	active.Sessions = blank
}

func GetSessionList() map[string]*types.Session {
	return sessionList
}

func NewSession(name string) string {
	yearOne := 1
	newSession := types.Session{&yearOne, name, &types.GlobalChronicle{types.WorldMap{}, []types.CultureYearGlobalChronicle{}}}
	if _, exists := sessionList[name]; !exists {
		sessionList[name] = &newSession
	} else {
		return "Name " + name + " is already in use!"
	}
	SetActiveSession(&newSession)
	return ""
}

func UpdateChronicle(worldMap types.WorldMap, chronica []types.CultureYearGlobalChronicle) {
	active.UpdateChronicle(worldMap, chronica)
}

func GetCurrentYear() int {
	return active.GetYear()
}

func EndThisYear() {
	active.EndYear()
}

func GetGlobalChronicle() types.GlobalChronicle {
	return active.GetChronicle()
}
