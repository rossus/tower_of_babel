package types

func (c SavedSubCulture) GetCode() []int {
	return c.Code
}
func (c SavedSubCulture) GetSubStage() int {
	return 0
}
func (c SavedSubCulture) GetStage() int {
	return 0
}
func (c SavedSubCulture) GetName() string {
	return ""
}

func (c SavedLocalCulture) GetCode() []int {
	return c.Code
}
func (c SavedLocalCulture) GetSubStage() int {
	return c.SubStage
}
func (c SavedLocalCulture) GetStage() int {
	return 0
}
func (c SavedLocalCulture) GetName() string {
	return ""
}

func (c SavedCulture) GetCode() []int {
	return c.Code
}
func (c SavedCulture) GetSubStage() int {
	return 0
}
func (c SavedCulture) GetStage() int {
	return c.Stage
}
func (c SavedCulture) GetName() string {
	return ""
}

func (c SavedBaseCulture) GetCode() []int {
	return c.Code
}
func (c SavedBaseCulture) GetSubStage() int {
	return 0
}
func (c SavedBaseCulture) GetStage() int {
	return 0
}
func (c SavedBaseCulture) GetName() string {
	return c.Name
}