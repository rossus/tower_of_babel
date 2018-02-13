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
	chronica:=chronicle.RunGlobalHistory(worldMap, 10)

	for i:=0; i<len(chronica[1].Cultures); i++ {
		fmt.Println(chronica[1].Cultures[i].Name, " ", chronica[1].Cultures[i].Stage, " (", chronica[1].Cultures[i].SubStage, ")")
	}
}
