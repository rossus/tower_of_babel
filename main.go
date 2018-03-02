package main

import (
	"github.com/rossus/tower_of_babel/chronicle"
	"github.com/rossus/tower_of_babel/maps"
	"fmt"
)

func main() {
	chronicle.WriteScriptorHeader()

	worldMap := maps.LoadAndConvertMap("mycenae2")
	chronica := chronicle.RunGlobalHistory(worldMap, 5000)

	//for i:=0; i<len(chronica[999].Cultures); i++ {
	//	fmt.Println(chronica[999].Cultures[i].Name, " ", chronica[999].Cultures[i].Stage, " (", chronica[999].Cultures[i].SubStage, ")")
	//}

	chronicle.WriteLocalCultureStages(chronica.WorldMap.Tiles[0][0].Chronica)

	fmt.Println(chronica.Chronica[0].Cultures[0].Area(chronica.WorldMap))
	fmt.Println(chronica.Chronica[0].Cultures[0].LocalCulture.Area(chronica.WorldMap))
	fmt.Println(chronica.Chronica[0].Cultures[0].Culture.Area(chronica.WorldMap))
	fmt.Println(chronica.Chronica[0].Cultures[0].BaseCulture.Area(chronica.WorldMap))

	//for i:=0; i<len(worldMap.Tiles[0][0].Chronica); i++ {
	//	fmt.Println(i+1, " --- ", chronica[0].Cultures[0].LocalCulture.YearArea(worldMap, i+1))
	//}

	fmt.Println(chronica.Chronica[0].Cultures[0].Started(chronica.Chronica), " - ", chronica.Chronica[0].Cultures[0].Ended(chronica.Chronica))
	fmt.Println(chronica.Chronica[0].Cultures[0].LocalCulture.Started(chronica.Chronica), " - ", chronica.Chronica[0].Cultures[0].LocalCulture.Ended(chronica.Chronica))
	fmt.Println(chronica.Chronica[0].Cultures[0].Culture.Started(chronica.Chronica), " - ", chronica.Chronica[0].Cultures[0].Culture.Ended(chronica.Chronica))
	fmt.Println(chronica.Chronica[0].Cultures[0].BaseCulture.Started(chronica.Chronica), " - ", chronica.Chronica[0].Cultures[0].BaseCulture.Ended(chronica.Chronica))

	chronicle.DrawYearCultureMap(worldMap.Tiles[0][0].Chronica[0].BaseCulture, 500, chronica)
	fmt.Println()
	chronicle.DrawYearCultureMap(worldMap.Tiles[0][0].Chronica[0].BaseCulture, 1000, chronica)
	fmt.Println()
	chronicle.DrawYearCultureMap(worldMap.Tiles[0][0].Chronica[0].BaseCulture, 2000, chronica)
	fmt.Println()
	chronicle.DrawYearCultureMap(worldMap.Tiles[0][0].Chronica[0].BaseCulture, 3000, chronica)
	fmt.Println()
	chronicle.DrawYearCultureMap(worldMap.Tiles[0][0].Chronica[0].BaseCulture, 4000, chronica)
	fmt.Println()
	chronicle.DrawYearCultureMap(worldMap.Tiles[0][0].Chronica[0].BaseCulture, 5000, chronica)
}
