package converter

import (
	"github.com/rossus/tower_of_babel/common/types"
	"io/ioutil"
	"encoding/json"
)

func newSession(name string) types.Session {
	yearOne := 1
	newSession := types.Session{&yearOne, name, &types.GlobalChronicle{types.WorldMap{}, []types.CultureYearGlobalChronicle{}}}
	return newSession
}

func changeCode(old types.CultureGeneCode, new []int) types.CultureGeneCode {
	for i := 0; i < len(new); i++ {
		if old[new[i]] == 0 {
			old[new[i]] = 1
		} else {
			old[new[i]] = 0
		}
	}
	return old
}

func loadLocalHistory(lastYear int, chronica []types.CultureYearLocalChronicle, loadedChronica []types.LoadedCultureYearLocalChronicle) []types.CultureYearLocalChronicle {
	for i := 2; i < lastYear; i++ {
		chronicon := chronica[i-2]
		chronicon.Year++
		chronicon.Event = 0
		for j := 0; j < len(loadedChronica); j++ {
			if loadedChronica[j].Year == i {
				chronicon.Event = loadedChronica[j].Event
				switch loadedChronica[j].Event {
				case 1:
					chronicon.Code = changeCode(chronicon.Code, loadedChronica[j].LoadedCulture.Code)
					break
				case 2:
					chronicon.Code = changeCode(chronicon.Code, loadedChronica[j].LoadedCulture.Code)
					chronicon.SubStage = loadedChronica[j].LoadedCulture.SubStage
					chronicon.LocalCulture.Code = chronicon.Code
					if chronicon.SubStage > *chronicon.LocalCulture.Last {
						*chronicon.LocalCulture.Last = chronicon.SubStage
					}
					break
				case 3:
					chronicon.Code = changeCode(chronicon.Code, loadedChronica[j].LoadedCulture.Code)
					chronicon.SubStage = 1
					lastSub := 1
					chronicon.LocalCulture.Last = &lastSub
					chronicon.Stage = loadedChronica[j].LoadedCulture.Stage
					chronicon.LocalCulture.Code = chronicon.Code
					chronicon.Culture.Code = chronicon.Code
					if chronicon.Stage > *chronicon.Culture.Last {
						*chronicon.Culture.Last = chronicon.Stage
					}
					break
				case 4:
					chronicon.Code = changeCode(chronicon.Code, loadedChronica[j].LoadedCulture.Code)
					chronicon.SubStage = 1
					chronicon.Stage = 1
					lastSub := 1
					last := 1
					chronicon.LocalCulture.Last = &lastSub
					chronicon.Culture.Last = &last
					chronicon.Name = loadedChronica[j].LoadedCulture.Name
					chronicon.LocalCulture.Code = chronicon.Code
					chronicon.Culture.Code = chronicon.Code
					chronicon.BaseCulture.Code = chronicon.Code
					break
				default:
				}
				break
			}
		}
		chronica = append(chronica, chronicon)
	}
	return chronica
}

func convertMap(loadedMap types.LoadedWorldMap, lastYear int) types.WorldMap {
	var worldMap types.WorldMap
	var zeroPoint types.CultureYearLocalChronicle
	for i := 0; i < len(loadedMap.Tiles); i++ {
		var tileline []types.Tile
		for j := 0; j < len(loadedMap.Tiles[i]); j++ {
			var newTile types.Tile
			var chronica []types.CultureYearLocalChronicle

			var chronicon types.CultureYearLocalChronicle
			if i == 0 && j == 0 {
				stageOneSub := 1
				stageOne := 1
				zeroPoint.Name = loadedMap.Tiles[i][j].InitialName
				zeroPoint.Stage = 1
				zeroPoint.SubStage = 1
				zeroPoint.Event = 0
				zeroPoint.Code = loadedMap.Tiles[i][j].InitialCode
				zeroPoint.LocalCulture.Code = loadedMap.Tiles[i][j].InitialCode
				zeroPoint.Culture.Code = loadedMap.Tiles[i][j].InitialCode
				zeroPoint.BaseCulture.Code = loadedMap.Tiles[i][j].InitialCode
				zeroPoint.Year = 1
				zeroPoint.Culture.Last = &stageOne
				zeroPoint.LocalCulture.Last = &stageOneSub
				chronicon = zeroPoint
			} else {
				chronicon = zeroPoint
			}
			chronica = append(chronica, chronicon)
			chronica = loadLocalHistory(lastYear, chronica, loadedMap.Tiles[i][j].Chronica)

			newTile.Geography = loadedMap.Tiles[i][j].Geography
			newTile.HasRiver = loadedMap.Tiles[i][j].HasRiver
			newTile.Chronica = chronica
			tileline = append(tileline, newTile)
		}
		worldMap.Tiles = append(worldMap.Tiles, tileline)
	}
	return worldMap
}

func addCultureToList(tile types.Tile, cultures []types.SubCulture, year int) []types.SubCulture {
	tileCulture := tile.Chronica[year].SubCulture
	for i := 0; i < len(cultures); i++ {
		if tileCulture == cultures[i] {
			return cultures
		}
	}
	return append(cultures, tileCulture)
}

func generateGlobalChronicle(session *types.Session) {
	chronicle := make([]types.CultureYearGlobalChronicle, *session.Year-1)
	for y := 0; y < *session.Year-1; y++ {
		var cultures = []types.SubCulture{}
		for i := 0; i < len(session.Chronicle.WorldMap.Tiles); i++ {
			for j := 0; j < len(session.Chronicle.WorldMap.Tiles[i]); j++ {
				cultures = addCultureToList(session.Chronicle.WorldMap.Tiles[i][j], cultures, y)
			}
		}
		chronicle[y] = types.CultureYearGlobalChronicle{y + 1, cultures}
	}
	session.Chronicle.Chronica = chronicle
}

func LoadSession(name string) (types.Session, string) {
	var loadedSession types.LoadedSession

	savedSessionJSON, err := ioutil.ReadFile("./saves/" + name + ".json")
	if err != nil {
		return types.Session{}, err.Error()
	}

	err = json.Unmarshal(savedSessionJSON, &loadedSession)
	if err != nil {
		return types.Session{}, err.Error()
	}

	newSession := newSession(name)
	newSession.Year = &loadedSession.Year
	newSession.Chronicle.WorldMap = convertMap(loadedSession.WorldMap, loadedSession.Year)
	
	generateGlobalChronicle(&newSession)

	return newSession, ""
}
