package culture

import (
	"github.com/rossus/tower_of_babel/common/types"
	"github.com/rossus/tower_of_babel/genom"
	"github.com/rossus/tower_of_babel/common/constants"
)

func checkCultureForEvolution(culture types.SubCulture, newCode types.CultureGeneCode) (types.SubCulture, int) {
	if genom.MatchDifference(culture.BaseCulture.Code, newCode) > constants.GENOM_CODE_BASE_CULTURE_RANGE-1 {
		return DevelopNewBaseCulture(culture.Name, newCode), 4
	} else if genom.MatchDifference(culture.Culture.Code, newCode) > constants.GENOM_CODE_CULTURE_RANGE-1 {
		return DevelopNewCulture(culture.Culture, newCode), 3
	} else if genom.MatchDifference(culture.LocalCulture.Code, newCode) > constants.GENOM_CODE_LOCAL_CULTURE_RANGE-1 {
		return DevelopNewLocalCulture(culture.LocalCulture, newCode), 2
	} else if genom.MatchDifference(culture.Code, newCode) > 0 {
		return DevelopNewSubCulture(culture, newCode), 1
	} else {
		return culture, 0
	}
}

func YearlyCultureMutation(culture types.SubCulture) (types.SubCulture, int) {
	newCode := genom.Mutate(culture.Code)
	return checkCultureForEvolution(culture, newCode)
}
