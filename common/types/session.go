package types

type Session struct {
	Year      *int
	Name      string
	Chronicle *GlobalChronicle
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