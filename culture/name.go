package culture

import "github.com/rossus/tower_of_babel/common/randomization"

func GenerateName() string {
	return randomization.MakeNewName()
}
