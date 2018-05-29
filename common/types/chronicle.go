package types

//Event codes: 1 - SubCulture changed, 2 - LocalCulture changed, 3 - Culture changed, 4 - BaseCulture changed
type CultureYearLocalChronicle struct {
	Year  int  `json:"year"`
	Event int  `json:"event"`
	SubCulture `json:"subculture"`
}

type CultureLocalTreeHistory []BaseCultureTreeHistory

type BaseCultureTreeHistory struct {
	BaseCulture                     `json:"base_cultures"`
	AlphaOmega [2]int               `json:"base_alpha_omega"`
	Cultures   []CultureTreeHistory `json:"cultures"`
}

type CultureTreeHistory struct {
	Culture                                 `json:"culture"`
	AlphaOmega    [2]int                    `json:"cul_alpha_omega"`
	LocalCultures []LocalCultureTreeHistory `json:"local_cultures"`
}

type LocalCultureTreeHistory struct {
	LocalCulture      `json:"local_culture"`
	AlphaOmega [2]int `json:"local_alpha_omega"`
}

type CultureYearGlobalChronicle struct {
	Year     int          `json:"year"`
	Cultures []SubCulture `json:"subcultures"`
}

type CultureGlobalTreeHistory []BaseCultureTreeHistory

type GlobalChronicle struct {
	WorldMap WorldMap                     `json:"world_map"`
	Chronica []CultureYearGlobalChronicle `json:"chronica"`
}

//Interface for all types of cultures from SubCulture to BaseCulture
type Cultures interface {
	YearArea(chronica GlobalChronicle, year int) int
	Area(chronica GlobalChronicle) (year []int, area int)
	Started(chronica GlobalChronicle) (year int)
	Ended(chronica GlobalChronicle) (year int)
	Type() string
}

type SavedCultureYearLocalChronicle struct {
	Year  int  `json:"year"`
	Event int  `json:"event"`
	SavedSubCulture `json:"culture"`
}