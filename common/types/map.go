package types

type Tile struct {
	Chronica []CultureYearLocalChronicle `json:"chronica"`
	HasRiver [4]bool	`json:"river"`									// is there a river on the tile border [NSWE]
	Geography string `json:"geography"`
}

type WorldMap struct {
	Tiles [][]Tile
}