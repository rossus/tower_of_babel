package chronicle

import (
	"github.com/rossus/tower_of_babel/common/types"
	"github.com/rossus/tower_of_babel/time"
)

func StartGlobalCultureChronicle(worldMap types.WorldMap, originCulture types.SubCulture, year int) []types.CultureYearGlobalChronicle {
	chronicle := make([]types.CultureYearGlobalChronicle, year)
	for i := 0; i < len(worldMap.Tiles); i++ {
		for j := 0; j < len(worldMap.Tiles[i]); j++ {
			worldMap.Tiles[i][j].Chronica = startLocalCultureChronicle(originCulture, year)
		}
	}
	chronicle[0] = types.CultureYearGlobalChronicle{1, []types.SubCulture{originCulture}}
	time.EndThisYear()
	return chronicle
}

func continueGlobalCultureChronicle(chronicle []types.CultureYearGlobalChronicle, cultures []types.SubCulture) []types.CultureYearGlobalChronicle {
	chronicle[time.GetCurrentYear()-1] = types.CultureYearGlobalChronicle{time.GetCurrentYear(), cultures}
	time.EndThisYear()
	return chronicle
}

func addCultureToList(tile types.Tile, cultures []types.SubCulture) []types.SubCulture {
	tileCulture:=tile.Chronica[time.GetCurrentYear()-1].SubCulture
	for i := 0; i < len(cultures); i++ {
		if tileCulture == cultures[i] {
			return cultures
		}
	}
	return append(cultures, tileCulture)
}

func expandLocalCultureChronicles(worldMap *types.WorldMap, addition int) {
	for i := 0; i < len(worldMap.Tiles); i++ {
		for j := 0; j < len(worldMap.Tiles[i]); j++ {
			worldMap.Tiles[i][j].Chronica = expandLocalCultureChronicle(worldMap.Tiles[i][j].Chronica, addition)
		}
	}
}