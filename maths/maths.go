package maths

import (
	"math/rand"
)

func randomBetweenInt(min, max int) int {
	return rand.Intn(max-min+1) + min
}

func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func maxInt(a, b, int) int {
	if a > b {
		return a
	}
	return b
}
