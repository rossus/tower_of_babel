package maps

import (
	"io/ioutil"
	"bytes"
	"strings"
	"github.com/rossus/tower_of_babel/common/types"
)



func loadMap(mapName string) ([][]byte, [][]byte, []string, []string){
	bs, err := ioutil.ReadFile("./maps/"+mapName)
	if err != nil {
		return [][]byte{}, [][]byte{}, []string{}, []string{}
	}

	var byteEnter byte = 10

	str := string(bs)

	byteBigMap:=bytes.Split(bs, []byte{byteEnter})
	stringBigMap:=strings.Split(str, string(byte(10)))

	byteSmallMap:=[][]byte{}
	stringSmallMap:=[]string{}

	for i:=0; i<len(byteBigMap); i++ {
		if i%2==1 {
			byteStroke:=[]byte{}
			stringStroke:=""
			for j:=0; j<len(byteBigMap[i]); j++ {
				if j%2==1 {
					byteStroke=append(byteStroke, byteBigMap[i][j])
					stringStroke=stringStroke+string(stringBigMap[i][j])
				}
			}
			byteSmallMap=append(byteSmallMap, byteStroke)
			stringSmallMap=append(stringSmallMap, stringStroke)
		}
	}

	return byteBigMap, byteSmallMap, stringBigMap, stringSmallMap
}

func LoadAndConvertMap(mapName string) types.WorldMap {
		bigMap, smallMap, _, _ := loadMap(mapName)
		worldMap:=types.WorldMap{}
		for i:=0; i<len(smallMap); i++ {
			worldMap.Tiles=append(worldMap.Tiles, []types.Tile{})
			for j:=0; j<len(smallMap[i]); j++ {
				tile:=types.Tile{}
				tile.Geography=string(smallMap[i][j])
				if bigMap[i*2][j*2+1]==byte('-') {tile.HasRiver[0]=true}
				if bigMap[i*2+2][j*2+1]==byte('-') {tile.HasRiver[1]=true}
				if bigMap[i*2+1][j*2]==byte('-') {tile.HasRiver[2]=true}
				if bigMap[i*2+1][j*2+2]==byte('-') {tile.HasRiver[3]=true}
				worldMap.Tiles[i]=append(worldMap.Tiles[i], tile)
			}
	}
	return worldMap
}
