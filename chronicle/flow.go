package chronicle

import (
	"github.com/rossus/tower_of_babel/culture"
	"github.com/rossus/tower_of_babel/time"
	"github.com/rossus/tower_of_babel/common/types"
)

func runLocalHistory(tile types.Tile) types.Tile {
		cultura:=tile.Chronica[time.GetCurrentYear()].SubCulture
		cultura, event:=culture.YearlyCultureMutation(cultura)
		tile.Chronica=ContinueLocalCultureChronicle(tile.Chronica, event, cultura)
	return tile
}

func RunGlobalHistory(worldMap types.WorldMap, finalYear int) []types.CultureYearGlobalChronicle {
	originCulture:=culture.MakeOriginalCulture()
	chronica:=StartGlobalCultureChronicle(worldMap, originCulture)
	for i:=0; i<finalYear-1; i++ {
		for j:=0; j<len(worldMap.Tiles); j++ {
			for k:=0 ; k<len(worldMap.Tiles[j]); k++ {
				worldMap.Tiles[j][k] = runLocalHistory(worldMap.Tiles[j][k])
			}
		}
		//chronica=ContinueLocalCultureChronicle(chronica, event, cultura)
		time.EndThisYear()
	}
	return chronica
}
