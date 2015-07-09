package average

import ()

type MovingAverage struct {
	SamplesCount int64
	Samples      []int64
	Size         int64
	Pointer      int64
}

func NewMovingAverage(size int64) *MovingAverage {
	return &MovingAverage{
		Size:         size,
		Samples:      make([]int64, 100),
		SamplesCount: 0,
		Pointer:      0,
	}
}

func (a *MovingAverage) Add(value int64) {
	a.Samples[a.Pointer%a.Size] = value
	a.Pointer++
	a.SamplesCount++
	if a.SamplesCount >= a.Size {
		a.SamplesCount = a.Size
	}
}

func (a *MovingAverage) Calculate() float64 {
	sum := int64(0)
	for _, i := range a.Samples {
		sum += i
	}

	return float64(sum) / float64(a.SamplesCount)
}
