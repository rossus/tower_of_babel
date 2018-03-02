package main

import (
	"github.com/rossus/tower_of_babel/chronicle"
	"github.com/rossus/tower_of_babel/maps"
	"fmt"
)

func main() {
	chronicle.WriteScriptorHeader()

	worldMap := maps.LoadAndConvertMap("mycenae2")
	chronica, worldMap := chronicle.RunGlobalHistory(worldMap, 5000)

	//for i:=0; i<len(chronica[999].Cultures); i++ {
	//	fmt.Println(chronica[999].Cultures[i].Name, " ", chronica[999].Cultures[i].Stage, " (", chronica[999].Cultures[i].SubStage, ")")
	//}

	chronicle.WriteLocalCultureStages(worldMap.Tiles[0][0].Chronica)

	fmt.Println(chronica[0].Cultures[0].Area(worldMap))
	fmt.Println(chronica[0].Cultures[0].LocalCulture.Area(worldMap))
	fmt.Println(chronica[0].Cultures[0].Culture.Area(worldMap))
	fmt.Println(chronica[0].Cultures[0].BaseCulture.Area(worldMap))

	//for i:=0; i<len(worldMap.Tiles[0][0].Chronica); i++ {
	//	fmt.Println(i+1, " --- ", chronica[0].Cultures[0].LocalCulture.YearArea(worldMap, i+1))
	//}

	fmt.Println(chronica[0].Cultures[0].Started(chronica), " - ", chronica[0].Cultures[0].Ended(chronica))
	fmt.Println(chronica[0].Cultures[0].LocalCulture.Started(chronica), " - ", chronica[0].Cultures[0].LocalCulture.Ended(chronica))
	fmt.Println(chronica[0].Cultures[0].Culture.Started(chronica), " - ", chronica[0].Cultures[0].Culture.Ended(chronica))
	fmt.Println(chronica[0].Cultures[0].BaseCulture.Started(chronica), " - ", chronica[0].Cultures[0].BaseCulture.Ended(chronica))

	chronicle.DrawYearCultureMap(worldMap.Tiles[0][0].Chronica[0].BaseCulture, 500, worldMap)
	fmt.Println()
	chronicle.DrawYearCultureMap(worldMap.Tiles[0][0].Chronica[0].BaseCulture, 1000, worldMap)
	fmt.Println()
	chronicle.DrawYearCultureMap(worldMap.Tiles[0][0].Chronica[0].BaseCulture, 2000, worldMap)
	fmt.Println()
	chronicle.DrawYearCultureMap(worldMap.Tiles[0][0].Chronica[0].BaseCulture, 3000, worldMap)
	fmt.Println()
	chronicle.DrawYearCultureMap(worldMap.Tiles[0][0].Chronica[0].BaseCulture, 4000, worldMap)
	fmt.Println()
	chronicle.DrawYearCultureMap(worldMap.Tiles[0][0].Chronica[0].BaseCulture, 5000, worldMap)
}
