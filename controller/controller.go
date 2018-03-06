package controller

import (
	"fmt"
	"bufio"
	"strings"
	"os"
	"strconv"
	"github.com/rossus/tower_of_babel/chronicle"
	"github.com/rossus/tower_of_babel/common/types"
)

func GreetAndAsk() int {
	fmt.Println()
	var finalYear int
	fmt.Println("Welcome to Tower of Babel v.0.03!")
	fmt.Print("Input the final year of history: ")
	fmt.Scanln(&finalYear)
	return finalYear
}

func Controller(chronica types.GlobalChronicle) {
	fmt.Println("You can now get more information about this world. What do you want to do next? Type 'help' to get the full list of commands.")
	for {
		var command string
		scanner:=bufio.NewScanner(os.Stdin)
		if scanner.Scan() {
			command=scanner.Text()
		}
		cmd:=strings.Split(command, " ")
		if cmd[0]=="atlas" {
			if len(cmd)>=3 {
				year, err := strconv.Atoi(cmd[2])
				if err != nil {
					fmt.Println(err)
				} else {
					chronicle.Atlas(chronica, cmd[1], year)
				}
			}
		} else if cmd[0]=="continue" {
			if len(cmd)>=2 {
				year, err := strconv.Atoi(cmd[1])
				if err != nil {
					fmt.Println(err)
				} else {
					chronicle.ContinueGlobalHistory(&chronica, year)
					fmt.Printf("This world has %v years of history now\n", len(chronica.Chronica))
					fmt.Println("")
				}
			}
		} else if cmd[0]=="exit" {
			fmt.Println("Bye!")
			break
		} else if cmd[0]=="help" {
			fmt.Println("You can use these commands:")
			fmt.Println("atlas [culture type] [year]		//draw an atlas for that type of original culture at certain year")
			fmt.Println("									culture types: s - subculture, l - local culture, c - culture, b - base culture")
			fmt.Println("continue [year]					//continue history up to this year	")
			fmt.Println("exit							//exit	")
			fmt.Println("help							//see all commands")
			fmt.Println("story							//tell a story about original culture	")
			fmt.Println("tree [x] [y]					//draw a culture tree for the tile (x, y)")
		} else if cmd[0]=="story" {
			chronicle.TellMeAStory(chronica)
		} else if cmd[0]=="tree" {
			if len(cmd)>=3 {
				x, err := strconv.Atoi(cmd[1])
				if err != nil {
					fmt.Println(err)
				} else {
					y, err := strconv.Atoi(cmd[2])
					if err != nil {
						fmt.Println(err)
					} else {
						chronicle.WriteLocalCultureStages(chronica.WorldMap.Tiles[y][x].Chronica)
					}
				}
			}
		}
	}
}