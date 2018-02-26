package main

import (
	"github.com/rossus/tower_of_babel/chronicle"
	"github.com/rossus/tower_of_babel/maps"
	"fmt"
)

func main() {
	chronicle.WriteScriptorHeader()
	//chronica:=chronicle.RunLocalHistory(5000)
	//chronicle.WriteLocalChronicle(chronica)
	//chronicle.WriteLocalCultureStages(chronica)

	worldMap:=maps.LoadAndConvertMap("mycenae2")
	chronica, worldMap:=chronicle.RunGlobalHistory(worldMap, 1000)

	for i:=0; i<len(chronica[999].Cultures); i++ {
		fmt.Println(chronica[999].Cultures[i].Name, " ", chronica[999].Cultures[i].Stage, " (", chronica[999].Cultures[i].SubStage, ")")
	}

	chronicle.WriteLocalCultureStages(worldMap.Tiles[0][0].Chronica)
}
