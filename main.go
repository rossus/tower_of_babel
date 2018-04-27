package main

import (
	"github.com/rossus/tower_of_babel/chronicle"
	"github.com/rossus/tower_of_babel/controller"
)

func main() {
	chronicle.WriteScriptorHeader()

	controller.MenuController()
}
