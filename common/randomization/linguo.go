package randomization


//getNewSyllable() returns random syllable.
func getNewSyllable() (syl string) {
	rnum := rnd().Intn(10)
	switch rnum {
	case 0:
		return "gwa"
	case 1:
		return "suul"
	case 2:
		return "ma"
	case 3:
		return "tos"
	case 4:
		return "ke"
	case 5:
		return "ar"
	case 6:
		return "que"
	case 7:
		return "run"
	case 8:
		return "gai"
	case 9:
		return "al"
	}
	return "aqua"
}

//decideSyllable(sylNum) will decide the fate of chosen syllable. If 0 - erase, if 1 - change, if 2 - add new syllable.
//If add new, then decide if it will be before(2), or after (3) chosen syllable. If word has 5 syllables, 0 and 1 are
//only avalible. If word has just one syllable, 0 is unavailible.
func decideSyllable(sylNum int) int {
	var fate int
	if sylNum >= 5 {
		fate = rnd().Intn(2)
	} else if sylNum == 1 {
		fate = rnd().Intn(2) + 1
	} else {
		fate = rnd().Intn(3)
	}
	if fate != 2 {
		return fate
	} else {
		return 2 + rnd().Intn(2)
	}
}

//matchSyllables(word) cuts a word into syllables. It will try to do it right at future.
func matchSyllables(name string) []string {
	var nameSyl []string
	for {
		length := len(name)
		if length < 4 {
			nameSyl = append(nameSyl, name)
			break
		} else {
			var horizon int
			if length < 6 {
				horizon = length - 2
			} else {
				horizon = 4
			}
			sylLen := rnd().Intn(horizon-1)+1
			nameSyl = append(nameSyl, name[0:sylLen])
			name = name[sylLen:length]
		}
	}
	return nameSyl
}

//MakeNewName() generates random name from random syllables.
func MakeNewName() (name string) {
	sylNum := rnd().Intn(5) + 1
	for i := 0; i < sylNum; i++ {
		name = name + getNewSyllable()
	}
	return name
}

//EvolveName(word) makes one word from another.
func EvolveName(name string) string {
	var newName string
	for {
		syllables := matchSyllables(name)
		chosenSyl := rnd().Intn(len(syllables))
		decision := decideSyllable(len(syllables))
		var newSyllables []string
		i := 0
		for i = 0; i < chosenSyl; i++ {
			newSyllables = append(newSyllables, syllables[i])
		}
		if decision == 0 {
			for j := chosenSyl; j < len(syllables)-1; j++ {
				newSyllables = append(newSyllables, syllables[j+1])
			}
		} else if decision == 1 {
			newSyllables = append(newSyllables, getNewSyllable())
			for j := chosenSyl + 1; j < len(syllables); j++ {
				newSyllables = append(newSyllables, syllables[j])
			}
		} else if decision == 2 {
			newSyllables = append(newSyllables, getNewSyllable())
			for j := chosenSyl; j < len(syllables); j++ {
				newSyllables = append(newSyllables, syllables[j])
			}
		} else if decision == 3 {
			newSyllables = append(newSyllables, syllables[chosenSyl])
			newSyllables = append(newSyllables, getNewSyllable())
			for j := chosenSyl; j < len(syllables); j++ {
				newSyllables = append(newSyllables, syllables[j])
			}
		}
		newName = ""
		for j := 0; j < len(newSyllables); j++ {
			newName = newName + newSyllables[j]
		}
		if newName != name {
			break
		}
	}
	return newName
}
