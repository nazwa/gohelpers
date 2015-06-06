package maths

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMinInt(t *testing.T) {
	assert.EqualValues(t, 10, MinInt(10, 20))
	assert.EqualValues(t, 10, MinInt(20, 10))

	assert.EqualValues(t, -10, MinInt(-10, 20))
	assert.EqualValues(t, -10, MinInt(20, -10))

	assert.EqualValues(t, -10, MinInt(-10, -5))
	assert.EqualValues(t, -10, MinInt(-5, -10))
}

func TestMaxInt(t *testing.T) {

	assert.EqualValues(t, 10, MaxInt(10, 5))
	assert.EqualValues(t, 10, MaxInt(5, 10))

	assert.EqualValues(t, 10, MaxInt(10, -5))
	assert.EqualValues(t, 10, MaxInt(-5, 10))

	assert.EqualValues(t, -10, MaxInt(-10, -15))
	assert.EqualValues(t, -10, MaxInt(-15, -10))

}

func TestRandBetweenInt(t *testing.T) {

	assert.EqualValues(t, 3, RandomBetweenInt(3, 3))
	assert.EqualValues(t, -3, RandomBetweenInt(-3, -3))
	assert.EqualValues(t, 0, RandomBetweenInt(0, 0))

}
