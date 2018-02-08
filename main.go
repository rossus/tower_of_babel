package main

import (
	"github.com/rossus/tower_of_babel/chronicle"
	"github.com/rossus/tower_of_babel/maps"
)

func main() {
	chronicle.WriteScriptorHeader()
	//chronica:=chronicle.RunLocalHistory(5000)
	//chronicle.WriteLocalChronicle(chronica)
	//chronicle.WriteLocalCultureStages(chronica)

	worldMap:=maps.LoadMap("mycenae2")
	chronica:=chronicle.RunGlobalHistory(worldMap, 5000)
}
