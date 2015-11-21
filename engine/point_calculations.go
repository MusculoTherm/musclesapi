package engine

import (
	"github.com/MusculoTherm/musclesapi/statistics"
)

func maxTempForPoint(arr []int64) int64 {
	var max int64 = arr[0]
	for _, temp := range arr {
		if temp > max {
			max = temp
		}
	}
	return max
}

func minTempForPoint(arr []int64) int64 {
	var min int64 = arr[0]
	for _, temp := range arr {
		if temp < min {
			min = temp
		}
	}
	return min
}

func q1TempForPoint(arr []int64) int64 {
	return statistics.GetQ1FromSortedInts(arr)
}

func q3TempForPoint(arr []int64) int64 {
	return statistics.GetQ3FromSortedInts(arr)
}

func medianTempForPoint(arr []int64) int64 {
	return statistics.MedianOfSortedInts(arr)
}

func meanTempForPoint(arr []int64) int64 {
	sum := 0.0
	for _, i := range arr {
		sum += float64(i)
	}
	return statistics.Round(sum / float64(len(arr)))
}

func iqrTempForPoint(arr []int64) int64 {
	return q3TempForPoint(arr) - q1TempForPoint(arr)
}
