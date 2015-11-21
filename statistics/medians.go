package statistics

import (
	"errors"
	"fmt"
	"sort"
)

func MedianOfSortedInts(orig []int64) int64 {
	input := I64ToF64(orig)
	l := len(input)
	if l%2 == 0 {
		tmp, _ := Mean(input[l/2-1 : l/2+1])
		return Round(tmp)
	} else {
		tmp := float64(input[l/2])
		return Round(tmp)
	}
}

func Median(input []float64) (median float64, err error) {
	c := copyslice(input)
	sort.Float64s(c)
	l := len(c)
	if l == 0 {
		return 0, errors.New("Input must not be empty")
	} else if l%2 == 0 {
		median, _ = Mean(c[l/2-1 : l/2+1])
	} else {
		median = float64(c[l/2])
	}

	return median, nil
}

func MatrixMedians(input [][]float64) ([]float64, error) {
	result := make([]float64, len(input[0]))
	for y, _ := range input[0] {
		var err error
		result[y], err = Median(getColumn(input, y))
		if err != nil {
			return result, errors.New(fmt.Sprintf("Error calculating medians for matrix: ", err))
		}
	}
	return result, nil
}
