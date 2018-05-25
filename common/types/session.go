package types

type Session struct {
	Year      *int `json:"year"`
	Name      string `json:"name"`
	Chronicle *GlobalChronicle `json:"chronicle"`
}

// Used when no session is chosen right now
type BlankSession struct{}

type ActiveSession struct {
	Sessions
}

//Common BlankSession and Session interface
type Sessions interface {
	GetYear() int
	EndYear()
	GetChronicle() GlobalChronicle
	UpdateChronicle(worldMap WorldMap, chronica []CultureYearGlobalChronicle)
}