package types

type Tile struct {
	Chronica []CultureYearLocalChronicle
	HasRiver [4]bool										// is there a river on the tile border [NSWE]
	Geography string
}

type WorldMap struct {
	Tiles [][]Tile
}