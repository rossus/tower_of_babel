package chronicle

import (
	"github.com/rossus/tower_of_babel/common/types"
	"github.com/rossus/tower_of_babel/time"
)

func startLocalCultureChronicle(culture types.SubCulture) []types.CultureYearLocalChronicle {
	chronicle := make([]types.CultureYearLocalChronicle, 0)
	chronicle=append(chronicle, types.CultureYearLocalChronicle{1, 0,culture})
	return chronicle
}

func continueLocalCultureChronicle(chronicle []types.CultureYearLocalChronicle, event int, culture types.SubCulture) []types.CultureYearLocalChronicle {
	chronica:=append(chronicle, types.CultureYearLocalChronicle{time.GetCurrentYear(), event, culture})
	return chronica
}

