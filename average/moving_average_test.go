package average

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSimpleMovingAverage(t *testing.T) {

	av := NewMovingAverage(10)
	av.Add(10)
	av.Add(10)
	assert.EqualValues(t, 10, av.Calculate())

}

func TestOverMovingAverage(t *testing.T) {

	av := NewMovingAverage(2)
	av.Add(10)
	av.Add(15)
	av.Add(20)
	assert.EqualValues(t, 17.5, av.Calculate())

}
