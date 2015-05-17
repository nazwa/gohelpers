package maths

import (
	"math/rand"
)

func RandomBetweenInt(min, max int) int {
	return rand.Intn(max-min+1) + min
}

func MinInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func MaxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}
