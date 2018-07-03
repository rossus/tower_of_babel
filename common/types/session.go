package types

type Session struct {
	Year      *int             `json:"year"`
	Name      string           `json:"name"`
	Chronicle *GlobalChronicle `json:"chronicle"`
}

// Used when no session is chosen right now
type BlankSession struct{}

type ActiveSession struct {
	Sessions
}

//Short info about existing session
type SessionInfo struct{
	Year int
	Name, Version string
}

//Common BlankSession and Session interface
type Sessions interface {
	GetYear() int
	EndYear()
	GetChronicle() GlobalChronicle
	UpdateChronicle(worldMap WorldMap, chronica []CultureYearGlobalChronicle)
	GetSession() Session
}

//Session struct used in saves
type SavedSession struct {
	Year     int           `json:"year"`
	Name     string        `json:"name"`
	Version  string        `json:"version"`
	WorldMap SavedWorldMap `json:"world_map"`
}

//Session struct used in loads
type LoadedSession struct {
	Year     int           `json:"year"`
	Name     string        `json:"name"`
	Version  string        `json:"version"`
	WorldMap LoadedWorldMap `json:"world_map"`
}
