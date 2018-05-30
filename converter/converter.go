package converter

import (
	"github.com/rossus/tower_of_babel/common/types"
	"os"
	"fmt"
	"encoding/json"
)

func SaveSession(s types.Session) {
	var saved types.SavedSession
	world_map := copyTilesFrom(s.Chronicle.WorldMap)
	saved.Year = *s.Year
	saved.Name = s.Name
	saved.WorldMap = world_map

	path := "./saves/" + s.Name + ".json"
	if _, err := os.Stat(path); !os.IsNotExist(err) {
		err := os.Remove(path)

		if err != nil {
			fmt.Println(err)
		}
	}
	file, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	sessionJSON, err := json.Marshal(saved)
	if err != nil {
		fmt.Println(err)
	}

	file.Write(sessionJSON)
}

func checkCodeChange(prevCode, code types.CultureGeneCode) []int {
	var newCode []int
	for i := 0; i < len(code); i++ {
		if prevCode[i] != code[i] {
			newCode = append(newCode, i)
		}
	}
	return newCode
}

func changeSubCulture(prevCode types.CultureGeneCode, yearChronica types.CultureYearLocalChronicle) types.SavedCultureYearLocalChronicle {
	var newChron types.SavedCultureYearLocalChronicle
	var c types.SavedSubCulture
	c.Code=checkCodeChange(prevCode, yearChronica.Code)
	newChron.Year = yearChronica.Year
	newChron.Event = yearChronica.Event
	newChron.SavedCultures = c
	return newChron
}

func changeLocalCulture(prevCode types.CultureGeneCode, yearChronica types.CultureYearLocalChronicle) types.SavedCultureYearLocalChronicle {
	var newChron types.SavedCultureYearLocalChronicle
	var c types.SavedLocalCulture
	c.SubStage=yearChronica.SubStage
	c.Code=checkCodeChange(prevCode, yearChronica.Code)
	newChron.Year = yearChronica.Year
	newChron.Event = yearChronica.Event
	newChron.SavedCultures = c
	return newChron
}

func changeCulture(prevCode types.CultureGeneCode, yearChronica types.CultureYearLocalChronicle) types.SavedCultureYearLocalChronicle {
	var newChron types.SavedCultureYearLocalChronicle
	var c types.SavedCulture
	c.Stage=yearChronica.Stage
	c.Code=checkCodeChange(prevCode, yearChronica.Code)
	newChron.Year = yearChronica.Year
	newChron.Event = yearChronica.Event
	newChron.SavedCultures = c
	return newChron
}

func changeBaseCulture(prevCode types.CultureGeneCode, yearChronica types.CultureYearLocalChronicle) types.SavedCultureYearLocalChronicle {
	var newChron types.SavedCultureYearLocalChronicle
	var c types.SavedBaseCulture
	c.Name=yearChronica.Name
	c.Code=checkCodeChange(prevCode, yearChronica.Code)
	newChron.Year = yearChronica.Year
	newChron.Event = yearChronica.Event
	newChron.SavedCultures = c
	return newChron
}

func checkChronica(chronica []types.CultureYearLocalChronicle) []types.SavedCultureYearLocalChronicle {
	var convChronica []types.SavedCultureYearLocalChronicle
	for i := 0; i < len(chronica); i++ {
		if i != 0 {
			switch chronica[i].Event {
			case 1:
				convChronica = append(convChronica, changeSubCulture(chronica[i-1].Code, chronica[i]))
				break
			case 2:
				convChronica = append(convChronica, changeLocalCulture(chronica[i-1].Code, chronica[i]))
				break
			case 3:
				convChronica = append(convChronica, changeCulture(chronica[i-1].Code, chronica[i]))
				break
			case 4:
				convChronica = append(convChronica, changeBaseCulture(chronica[i-1].Code, chronica[i]))
				break
			default:
			}
		}
	}
	return convChronica
}

func copyTilesFrom(m types.WorldMap) types.SavedWorldMap {
	var world_map types.SavedWorldMap
	for i := 0; i < len(m.Tiles); i++ {
		var tileline []types.SavedTile
		for j := 0; j < len(m.Tiles[i]); j++ {
			var newTile types.SavedTile
			newTile.Geography = m.Tiles[i][j].Geography
			newTile.HasRiver = m.Tiles[i][j].HasRiver
			newTile.InitialCode = m.Tiles[i][j].Chronica[0].Code
			newTile.InitialName = m.Tiles[i][j].Chronica[0].Name
			newTile.Chronica = checkChronica(m.Tiles[i][j].Chronica)
			tileline = append(tileline, newTile)
		}
		world_map.Tiles = append(world_map.Tiles, tileline)
	}
	return world_map
}
