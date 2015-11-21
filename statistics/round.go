package statistics

import (
	"math"
)

func Round(f float64) int64 {
	return int64(math.Floor(f + .5))
}

func RoundArray(f []float64) []int64 {
	out := make([]int64, len(f))
	for ind, item := range f {
		out[ind] = Round(item)
	}
	return out
}
