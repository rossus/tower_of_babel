package types

import "github.com/rossus/tower_of_babel/common/constants"

type CultureGeneCode [constants.GENOM_CODE_LENGTH]byte

//type MicroGeneCode int
//
//type CultureGene struct {
//	Language MicroGeneCode
//	Fashion MicroGeneCode
//	Architecture MicroGeneCode
//}

type BaseCulture struct {
	Name string `json:"name"`
	Code CultureGeneCode `json:"code"`
	//Ancestor AncestorCulture
}

type Culture struct {
	BaseCulture
	Stage int `json:"stage"`
	Last *int	`json:"last"`			//youngest stage of this culture in the world
	Code CultureGeneCode `json:"code"`
}

type LocalCulture struct {
	Culture
	SubStage int `json:"substage"`
	Last *int	`json:"last_sub"`			//youngest stage of this culture in the world
	Code CultureGeneCode `json:"code"`
}

type SubCulture struct {
	LocalCulture
	Code CultureGeneCode `json:"code"`
}