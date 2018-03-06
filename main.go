package main

import (
	"github.com/rossus/tower_of_babel/chronicle"
	"github.com/rossus/tower_of_babel/maps"
	"github.com/rossus/tower_of_babel/controller"
)

func main() {
	chronicle.WriteScriptorHeader()

	worldMap := maps.LoadAndConvertMap("mycenae2")

	finalYear:=controller.GreetAndAsk()

	chronica := chronicle.RunGlobalHistory(worldMap, finalYear)

	chronicle.TellMeAStory(chronica)

	controller.Controller(chronica)
}
