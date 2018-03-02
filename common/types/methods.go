package types

//YearArea methods
func (culture BaseCulture) YearArea(chronica GlobalChronicle, year int) int {
	worldMap:=chronica.WorldMap
	area := 0
	for i := 0; i < len(worldMap.Tiles); i++ {
		for j := 0; j < len(worldMap.Tiles[i]); j++ {
			if worldMap.Tiles[i][j].Chronica[year-1].BaseCulture == culture {
				area++
			}
		}
	}
	return area
}

func (culture Culture) YearArea(chronica GlobalChronicle, year int) int {
	worldMap:=chronica.WorldMap
	area := 0
	for i := 0; i < len(worldMap.Tiles); i++ {
		for j := 0; j < len(worldMap.Tiles[i]); j++ {
			if worldMap.Tiles[i][j].Chronica[year-1].Culture == culture {
				area++
			}
		}
	}
	return area
}

func (culture LocalCulture) YearArea(chronica GlobalChronicle, year int) int {
	worldMap:=chronica.WorldMap
	area := 0
	for i := 0; i < len(worldMap.Tiles); i++ {
		for j := 0; j < len(worldMap.Tiles[i]); j++ {
			if worldMap.Tiles[i][j].Chronica[year-1].LocalCulture == culture {
				area++
			}
		}
	}
	return area
}

func (culture SubCulture) YearArea(chronica GlobalChronicle, year int) int {
	worldMap:=chronica.WorldMap
	area := 0
	for i := 0; i < len(worldMap.Tiles); i++ {
		for j := 0; j < len(worldMap.Tiles[i]); j++ {
			if worldMap.Tiles[i][j].Chronica[year-1].SubCulture == culture {
				area++
			}
		}
	}
	return area
}

//Area methods
func (culture BaseCulture) Area(chronica GlobalChronicle) (year []int, area int) {
	worldMap:=chronica.WorldMap
	area = 0
	year = []int{}
	for i := 0; i < len(worldMap.Tiles[0][0].Chronica); i++ {
		yearArea := culture.YearArea(chronica, i+1)
		if yearArea == area {
			year = append(year, i+1)
		} else if yearArea > area {
			year = []int{}
			area = yearArea
			year = append(year, i+1)
		}
	}
	return
}

func (culture Culture) Area(chronica GlobalChronicle) (year []int, area int) {
	worldMap:=chronica.WorldMap
	area = 0
	year = []int{}
	for i := 0; i < len(worldMap.Tiles[0][0].Chronica); i++ {
		yearArea := culture.YearArea(chronica, i+1)
		if yearArea == area {
			year = append(year, i+1)
		} else if yearArea > area {
			year = []int{}
			area = yearArea
			year = append(year, i+1)
		}
	}
	return
}

func (culture LocalCulture) Area(chronica GlobalChronicle) (year []int, area int) {
	worldMap:=chronica.WorldMap
	area = 0
	year = []int{}
	for i := 0; i < len(worldMap.Tiles[0][0].Chronica); i++ {
		yearArea := culture.YearArea(chronica, i+1)
		if yearArea == area {
			year = append(year, i+1)
		} else if yearArea > area {
			year = []int{}
			area = yearArea
			year = append(year, i+1)
		}
	}
	return
}

func (culture SubCulture) Area(chronica GlobalChronicle) (year []int, area int) {
	worldMap:=chronica.WorldMap
	area = 0
	year = []int{}
	for i := 0; i < len(worldMap.Tiles[0][0].Chronica); i++ {
		yearArea := culture.YearArea(chronica, i+1)
		if yearArea == area {
			year = append(year, i+1)
		} else if yearArea > area {
			year = []int{}
			area = yearArea
			year = append(year, i+1)
		}
	}
	return
}

//Started methods
func (culture BaseCulture) Started(chronica GlobalChronicle) (year int) {
	chronicle:=chronica.Chronica
	for i := 0; i < len(chronicle); i++ {
		for j := 0; j < len(chronicle[i].Cultures); j++ {
			if culture == chronicle[i].Cultures[j].BaseCulture {
				return i + 1
			}
		}
	}
	return
}

func (culture Culture) Started(chronica GlobalChronicle) (year int) {
	chronicle:=chronica.Chronica
	for i := 0; i < len(chronicle); i++ {
		for j := 0; j < len(chronicle[i].Cultures); j++ {
			if culture == chronicle[i].Cultures[j].Culture {
				return i + 1
			}
		}
	}
	return
}

func (culture LocalCulture) Started(chronica GlobalChronicle) (year int) {
	chronicle:=chronica.Chronica
	for i := 0; i < len(chronicle); i++ {
		for j := 0; j < len(chronicle[i].Cultures); j++ {
			if culture == chronicle[i].Cultures[j].LocalCulture {
				return i + 1
			}
		}
	}
	return
}

func (culture SubCulture) Started(chronica GlobalChronicle) (year int) {
	chronicle:=chronica.Chronica
	for i := 0; i < len(chronicle); i++ {
		for j := 0; j < len(chronicle[i].Cultures); j++ {
			if culture == chronicle[i].Cultures[j] {
				return i + 1
			}
		}
	}
	return
}

//Ended methods
func (culture BaseCulture) Ended(chronica GlobalChronicle) (year int) {
	chronicle:=chronica.Chronica
	for i := culture.Started(chronica) - 1; i < len(chronicle); i++ {
		sign := false
		for j := 0; j < len(chronicle[i].Cultures); j++ {
			if culture == chronicle[i].Cultures[j].BaseCulture {
				sign = true
			}
		}
		if sign == false {
			return i
		}
	}
	return 0
}

func (culture Culture) Ended(chronica GlobalChronicle) (year int) {
	chronicle:=chronica.Chronica
	for i := culture.Started(chronica) - 1; i < len(chronicle); i++ {
		sign := false
		for j := 0; j < len(chronicle[i].Cultures); j++ {
			if culture == chronicle[i].Cultures[j].Culture {
				sign = true
			}
		}
		if sign == false {
			return i
		}
	}
	return 0
}

func (culture LocalCulture) Ended(chronica GlobalChronicle) (year int) {
	chronicle:=chronica.Chronica
	for i := culture.Started(chronica) - 1; i < len(chronicle); i++ {
		sign := false
		for j := 0; j < len(chronicle[i].Cultures); j++ {
			if culture == chronicle[i].Cultures[j].LocalCulture {
				sign = true
			}
		}
		if sign == false {
			return i
		}
	}
	return 0
}

func (culture SubCulture) Ended(chronica GlobalChronicle) (year int) {
	chronicle:=chronica.Chronica
	for i := culture.Started(chronica) - 1; i < len(chronicle); i++ {
		sign := false
		for j := 0; j < len(chronicle[i].Cultures); j++ {
			if culture == chronicle[i].Cultures[j] {
				sign = true
			}
		}
		if sign == false {
			return i
		}
	}
	return 0
}

//Type methods
func (culture BaseCulture) Type() string {
	return "BaseCulture"
}

func (culture Culture) Type() string {
	return "Culture"
}

func (culture LocalCulture) Type() string {
	return "LocalCulture"
}

func (culture SubCulture) Type() string {
	return "SubCulture"
}