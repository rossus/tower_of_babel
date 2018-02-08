package chronicle

import (
	"github.com/rossus/tower_of_babel/common/types"
	"github.com/rossus/tower_of_babel/time"
)

//TODO: write area covered by culture
func StartGlobalCultureChronicle(worldMap types.WorldMap, originCulture types.SubCulture) []types.CultureYearGlobalChronicle {
	chronicle := make([]types.CultureYearGlobalChronicle, 0)
	for i := 0; i < len(worldMap.Tiles); i++ {
		for j := 0; j < len(worldMap.Tiles[i]); j++ {
			worldMap.Tiles[i][j].Chronica = startLocalCultureChronicle(originCulture)
		}
	}
	chronicle = append(chronicle, types.CultureYearGlobalChronicle{1, []types.SubCulture{originCulture}})
	time.EndThisYear()
	return chronicle
}

func ContinueGlobalCultureChronicle(chronicle []types.CultureYearGlobalChronicle, cultures []types.SubCulture) []types.CultureYearGlobalChronicle {
	chronica:=append(chronicle, types.CultureYearGlobalChronicle{time.GetCurrentYear(), cultures})
	time.EndThisYear()
	return chronica
}

func AddCultureToList(tile types.Tile, cultures []types.SubCulture) []types.SubCulture {
	tileCulture:=tile.Chronica[time.GetCurrentYear()].SubCulture
	for i := 0; i < len(cultures); i++ {
		if tileCulture == cultures[i] {
			return cultures
		}
	}
	return append(cultures, tileCulture)
}
