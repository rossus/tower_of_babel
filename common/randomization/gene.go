package randomization

import "github.com/rossus/tower_of_babel/common/constants"

func RandByte() (b byte) {
	rnum := rnd().Intn(2)
	if rnum == 0 {
		return 0
	} else {
		return 1
	}
}

func GenomMutate() (b byte, i int) {
	rnum := rnd().Intn(constants.CULTURE_TOTAL_CHANCE_PUL)
	if rnum <= constants.CULTURE_MUTATION_RATE {
		i = rnd().Intn(constants.GENOM_CODE_LENGTH-1)
		rnum = rnd().Intn(2)
		if rnum == 0 {
			return 0, i
		} else {
			return 1, i
		}
	} else {
		return 0, -1
	}
}