package randomization

import (
	"time"
	"math/rand"
)

func rnd() (rnd *rand.Rand) {
	src := rand.NewSource(time.Now().UnixNano())
	return rand.New(src)
}
