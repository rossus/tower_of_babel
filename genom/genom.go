package genom

import (
	"github.com/rossus/tower_of_babel/common/types"
	"github.com/rossus/tower_of_babel/common/randomization"
	"github.com/rossus/tower_of_babel/common/constants"
)

func Generate() types.CultureGeneCode {
	var baseCultureCode types.CultureGeneCode
	for i := 0; i < len(baseCultureCode); i++ {
		baseCultureCode[i] = randomization.RandByte()
	}
	return baseCultureCode
}

func Mutate(code types.CultureGeneCode) types.CultureGeneCode {
	decision, place := randomization.GenomMutate()
	if place == -1 {
		return code
	} else {
		newCode := code
		newCode[place] = decision
		return newCode
	}
}

func MatchDifference(initial, object types.CultureGeneCode) int {
	difference:=0
	for i := 0; i < constants.GENOM_CODE_LENGTH; i++ {
		if initial[i] != object[i] {difference++}
	}
	return difference
}
