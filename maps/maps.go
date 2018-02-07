package maps

import (
	"io/ioutil"
	"bytes"
	"strings"
)



func LoadMap(mapName string) ([][]byte, [][]byte, []string, []string){
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

	return byteBigMap, byteSmallMap, stringBigMap, stringBigMap
}
