package types

//Methods for Sessions interface

func (s Session) GetYear() int {
	return *s.Year
}

func (s BlankSession) GetYear() int {
	return 0
}

func (s Session) EndYear() {
	*s.Year+=1
}

func (s BlankSession) EndYear() {}

func (s Session) GetChronicle() GlobalChronicle {
	return *s.Chronicle
}

func (s BlankSession) GetChronicle() GlobalChronicle {
	return GlobalChronicle{}
}

func (s Session) UpdateChronicle(worldMap WorldMap, chronica []CultureYearGlobalChronicle) {
	*s.Chronicle=GlobalChronicle{worldMap, chronica}
}

func (s BlankSession) UpdateChronicle(worldMap WorldMap, chronica []CultureYearGlobalChronicle) {}