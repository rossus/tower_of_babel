package main

import (
	"github.com/rossus/tower_of_babel/chronicle"
	"github.com/rossus/tower_of_babel/maps"
)

func main() {
	chronicle.WriteScriptorHeader()

	worldMap := maps.LoadAndConvertMap("mycenae2")
	chronica := chronicle.RunGlobalHistory(worldMap, 5000)

	//chronicle.WriteLocalCultureStages(chronica.WorldMap.Tiles[0][0].Chronica)

	chronicle.Atlas(chronica)

	chronicle.TellMeAStory(chronica)
}
