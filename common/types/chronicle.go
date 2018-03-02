package types

//Event codes: 1 - SubCulture changed, 2 - LocalCulture changed, 3 - Culture changed, 4 - BaseCulture changed
type CultureYearLocalChronicle struct {
	Year int
	Event int
	SubCulture
}

type CultureLocalTreeHistory []BaseCultureTreeHistory

type BaseCultureTreeHistory struct {
	BaseCulture
	AlphaOmega [2]int
	Cultures []CultureTreeHistory
}

type CultureTreeHistory struct {
	Culture
	AlphaOmega [2]int
	LocalCultures []LocalCultureTreeHistory
}

type LocalCultureTreeHistory struct {
	LocalCulture
	AlphaOmega [2]int
}

type CultureYearGlobalChronicle struct {
	Year int
	Cultures []SubCulture
}

type CultureGlobalTreeHistory []BaseCultureTreeHistory

type GlobalChronicle struct {
	WorldMap WorldMap
	Chronica []CultureYearGlobalChronicle
}

//Interface for all types of cultures from SubCulture to BaseCulture
type Cultures interface {
	YearArea(chronica GlobalChronicle, year int) int
	Area(chronica GlobalChronicle) (year []int, area int)
	Started(chronica GlobalChronicle) (year int)
	Ended(chronica GlobalChronicle) (year int)
	Type() string
}