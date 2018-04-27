package chronicle

import (
	"github.com/rossus/tower_of_babel/common/types"
	"github.com/rossus/tower_of_babel/session"
)

func startLocalCultureChronicle(culture types.SubCulture, year int) []types.CultureYearLocalChronicle {
	chronicle := make([]types.CultureYearLocalChronicle, year)
	chronicle[0] = types.CultureYearLocalChronicle{1, 0, culture}
	return chronicle
}

func continueLocalCultureChronicle(chronicle []types.CultureYearLocalChronicle, event int, culture types.SubCulture) []types.CultureYearLocalChronicle {
	chronicle[session.GetCurrentYear()-1] = types.CultureYearLocalChronicle{session.GetCurrentYear(), event, culture}
	return chronicle
}

func expandLocalCultureChronicle(chronicle []types.CultureYearLocalChronicle, addition int) []types.CultureYearLocalChronicle {
	additionalChronicle := make([]types.CultureYearLocalChronicle, addition)
	chronicle = append(chronicle, additionalChronicle...)
	return chronicle
}