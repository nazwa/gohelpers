package gohelpers

import (
	"math/rand"
	"time"
)

func RandomNumber(min, max int64, seed bool) int64 {

	if seed {
		rand.Seed(time.Now().UTC().UnixNano())
	}

	return rand.Int63n(max-min) + min
}
