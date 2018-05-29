package types

type Tile struct {
	Chronica  []CultureYearLocalChronicle `json:"chronica"`
	HasRiver  [4]bool                     `json:"river"` // is there a river on the tile border [NSWE]
	Geography string                      `json:"geography"`
}

type WorldMap struct {
	Tiles [][]Tile
}

//Map structs used in saves
type SavedTile struct {
	Chronica    []SavedCultureYearLocalChronicle `json:"chronica"`
	HasRiver    [4]bool                          `json:"river"` // is there a river on the tile border [NSWE]
	Geography   string                           `json:"geography"`
	InitialCode CultureGeneCode                  `json:"init_code"`
	InitialName string                           `json:"init_name"`
}

type SavedWorldMap struct {
	Tiles [][]SavedTile
}
