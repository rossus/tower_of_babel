package controller

import (
	"fmt"
	"bufio"
	"strings"
	"os"
	"strconv"
	"github.com/rossus/tower_of_babel/chronicle"
	"github.com/rossus/tower_of_babel/session"
	"github.com/rossus/tower_of_babel/maps"
)

func SessionController() {
	fmt.Println("You can now get more information about this world. What do you want to do next? Type 'help' to get the full list of commands.")
	for {
		var command string
		chronica:=session.GetGlobalChronicle()
		scanner := bufio.NewScanner(os.Stdin)
		if scanner.Scan() {
			command = scanner.Text()
		}
		cmd := strings.Split(command, " ")
		if cmd[0] == "atlas" {
			if len(cmd) >= 3 {
				year, err := strconv.Atoi(cmd[2])
				if err != nil {
					fmt.Println(err)
				} else {
					if year <= len(chronica.Chronica) {
						chronicle.Atlas(chronica, cmd[1], year)
					}
				}
			}
		} else if cmd[0] == "continue" {
			if len(cmd) >= 2 {
				year, err := strconv.Atoi(cmd[1])
				if err != nil {
					fmt.Println(err)
				} else if year >= session.GetCurrentYear() {
					chronicle.ContinueGlobalHistory(year)
					fmt.Printf("This world has %v years of history now\n", len(chronica.Chronica))
					fmt.Println("")
				}
			}
		} else if cmd[0] == "exit" {
			fmt.Println("Leaving the session...")
			session.UnloadSession()
			fmt.Println("------")
			break
		} else if cmd[0] == "help" {
			fmt.Println("You can use these commands:")
			fmt.Println("atlas [culture type] [year]		//draw an atlas for that type of original culture at certain year")
			fmt.Println("									culture types: s - subculture, l - local culture, c - culture, b - base culture")
			fmt.Println("continue [year]					//continue history up to this year	")
			fmt.Println("exit							//exit to menu")
			fmt.Println("help							//see all commands")
			fmt.Println("story							//tell a story about original culture	")
			fmt.Println("tree [x] [y]					//draw a culture tree for the tile (x, y)")
		} else if cmd[0] == "story" {
			chronicle.TellMeAStory(chronica)
		} else if cmd[0] == "tree" {
			if len(cmd) >= 3 {
				x, err := strconv.Atoi(cmd[1])
				if err != nil {
					fmt.Println(err)
				} else if (x <= len(chronica.WorldMap.Tiles[0]) && x >= 0) {
					y, err := strconv.Atoi(cmd[2])
					if err != nil {
						fmt.Println(err)
					} else if (y <= len(chronica.WorldMap.Tiles) && y >= 0) {
						chronicle.WriteLocalCultureStages(chronica.WorldMap.Tiles[y][x].Chronica)
					}
				}
			}
		}
	}
}

func MenuController() {
	fmt.Println()
	fmt.Println("Welcome to Tower of Babel v.0.03!")
	fmt.Println("What do you want to do? Type 'help' to get the full list of commands.")
	for {
		var command string
		scanner := bufio.NewScanner(os.Stdin)
		if scanner.Scan() {
			command = scanner.Text()
		}
		cmd := strings.Split(command, " ")
		if cmd[0] == "exit" {
			fmt.Println("Bye!")
			break
		} else if cmd[0] == "help" {
			fmt.Println("You can use these commands:")
			fmt.Println("exit							//exit	")
			fmt.Println("help							//see all commands")
			fmt.Println("list							//see all sessions")
			fmt.Println("load [session name]				//load session")
			fmt.Println("new [session name]				//start new session")
		} else if cmd[0] == "list" {
			list := session.GetSessionList()
			for x := range list {
				fmt.Println(" --- Session ", x, " : ", list[x].GetYear(), " years")
			}
		} else if cmd[0] == "load" {
			list := session.GetSessionList()
			if len(cmd) >= 2 {
				if _, exists := list[cmd[1]]; !exists {
					fmt.Println("No such session: ", cmd[1], "")
				} else {
					fmt.Println("------")
					session.SetActiveSession(list[cmd[1]])
					SessionController()
				}
			}
		} else if cmd[0] == "new" {
			list := session.GetSessionList()
			var finalYear int
			if len(cmd) >= 2 {
				if _, exists := list[cmd[1]]; exists {
					fmt.Println("Name " + cmd[1] + " is already in use!")
				} else {
					err := session.NewSession(cmd[1])
					if err != "" {
						fmt.Println(err)
					} else {
						fmt.Println("------")
						worldMap := maps.LoadAndConvertMap("mycenae2")
						fmt.Print("Enter the final year of history: ")
						fmt.Scanln(&finalYear)

						chronicle.RunGlobalHistory(worldMap, finalYear)

						chronicle.TellMeAStory(session.GetGlobalChronicle())

						SessionController()

					}

				}
			}
		}
	}
}
