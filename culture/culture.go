package culture

import (
	"github.com/rossus/tower_of_babel/common/types"
	"github.com/rossus/tower_of_babel/genom"
	"github.com/rossus/tower_of_babel/common/randomization"
)

func makeBaseCulture() types.BaseCulture {
	var culture types.BaseCulture
	culture.Code = genom.Generate()
	culture.Name = GenerateName()
	return culture
}

func makeCulture(baseCulture types.BaseCulture) types.Culture {
	var culture types.Culture
	var newLast = 1
	culture.BaseCulture = baseCulture
	culture.Code = baseCulture.Code
	culture.Stage = 1
	culture.Last = &newLast
	return culture
}

func makeLocalCulture(stageCulture types.Culture) types.LocalCulture {
	var culture types.LocalCulture
	var newLast = 1
	culture.Culture = stageCulture
	culture.Code = stageCulture.Code
	culture.SubStage = 1
	culture.Last = &newLast
	return culture
}

func makeSubCulture(localCulture types.LocalCulture) types.SubCulture {
	var culture types.SubCulture
	culture.LocalCulture = localCulture
	culture.Code = localCulture.Code
	return culture
}

func MakeOriginalCulture() types.SubCulture {
	return makeSubCulture(makeLocalCulture(makeCulture(makeBaseCulture())))
}

func DevelopNewBaseCulture(name string, code types.CultureGeneCode) types.SubCulture {
	var culture types.BaseCulture
	culture.Code = code
	culture.Name = randomization.EvolveName(name)
	return makeSubCulture(makeLocalCulture(makeCulture(culture)))
}

func DevelopNewCulture(init types.Culture, code types.CultureGeneCode) types.SubCulture {
	var culture types.Culture
	culture.Code = code
	culture.BaseCulture = init.BaseCulture
	culture.Last = init.Last
	*culture.Last++
	culture.Stage = *culture.Last
	return makeSubCulture(makeLocalCulture(culture))
}

func DevelopNewLocalCulture(init types.LocalCulture, code types.CultureGeneCode) types.SubCulture {
	var culture types.LocalCulture
	culture.Code = code
	culture.Culture = init.Culture
	culture.Last = init.Last
	*culture.Last++
	culture.SubStage = *culture.Last
	return makeSubCulture(culture)
}

func DevelopNewSubCulture(init types.SubCulture, code types.CultureGeneCode) types.SubCulture {
	var culture types.SubCulture
	culture.Code = code
	culture.LocalCulture = init.LocalCulture
	return culture
}
