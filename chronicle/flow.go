package chronicle

import (
	"github.com/rossus/tower_of_babel/culture"
	"github.com/rossus/tower_of_babel/session"
	"github.com/rossus/tower_of_babel/common/types"
)

func runLocalHistory(tile types.Tile) types.Tile {
	cultura := tile.Chronica[session.GetCurrentYear()-2].SubCulture
	cultura, event := culture.YearlyCultureMutation(cultura)
	tile.Chronica = continueLocalCultureChronicle(tile.Chronica, event, cultura)
	return tile
}

func RunGlobalHistory(worldMap types.WorldMap, finalYear int) {
	originCulture := culture.MakeOriginalCulture()
	chronica := StartGlobalCultureChronicle(worldMap, originCulture, finalYear)
	for i := 0; i < finalYear-1; i++ {
		var cultures = []types.SubCulture{}
		for j := 0; j < len(worldMap.Tiles); j++ {
			for k := 0; k < len(worldMap.Tiles[j]); k++ {
				worldMap.Tiles[j][k] = runLocalHistory(worldMap.Tiles[j][k])
				cultures = addCultureToList(worldMap.Tiles[j][k], cultures)
			}
		}
		continueGlobalCultureChronicle(chronica, cultures)
	}

	session.UpdateChronicle(worldMap, chronica)

}

func ContinueGlobalHistory(newFinal int) {
	chronica:=session.GetGlobalChronicle()
	oldAge := len(chronica.Chronica)
	newAge := newFinal-oldAge
	newAgeChronica := make([]types.CultureYearGlobalChronicle, newAge)
	chronica.Chronica=append(chronica.Chronica, newAgeChronica...)
	expandLocalCultureChronicles(&chronica.WorldMap, newAge)
	for i := oldAge; i < newFinal; i++ {
		var cultures = []types.SubCulture{}
		for j := 0; j < len(chronica.WorldMap.Tiles); j++ {
			for k := 0; k < len(chronica.WorldMap.Tiles[j]); k++ {
				chronica.WorldMap.Tiles[j][k] = runLocalHistory(chronica.WorldMap.Tiles[j][k])
				cultures = addCultureToList(chronica.WorldMap.Tiles[j][k], cultures)
			}
		}
		chronica.Chronica = continueGlobalCultureChronicle(chronica.Chronica, cultures)
	}
	session.UpdateChronicle(chronica.WorldMap, chronica.Chronica)
}