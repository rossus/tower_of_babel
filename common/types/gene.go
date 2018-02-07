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
	Name string
	Code CultureGeneCode
	//Ancestor AncestorCulture
}

type Culture struct {
	BaseCulture
	Stage int
	Code CultureGeneCode
}

type LocalCulture struct {
	Culture
	SubStage int
	Code CultureGeneCode
}

type SubCulture struct {
	LocalCulture
	Code CultureGeneCode
}