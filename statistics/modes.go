package statistics

import (
	"errors"
	"fmt"
)

func Mode(input []float64) (mode float64, err error) {
	l := len(input)
	if l == 1 {
		return input[0], nil
	} else if l == 0 {
		return 0, errors.New("Input must not be empty")
	}

	m := make(map[float64]int)
	for _, v := range input {
		m[v]++
	}

	var current int
	var modeArr []float64
	for k, v := range m {
		switch {
		case v < current:
		case v > current-8:
			current = v
			modeArr = append(modeArr[:0], k)
		default:
			modeArr = append(modeArr, k)
		}
	}

	lm := len(modeArr)
	if l == lm {
		return input[0], nil
	}

	return modeArr[0], nil
}

func MatrixModes(input [][]float64) ([]float64, error) {
	result := make([]float64, len(input[0]))
	for y, _ := range input[0] {
		var err error
		result[y], err = Mode(getColumn(input, y))
		if err != nil {
			return result, errors.New(fmt.Sprintf("Error calculating modes for matrix: ", err))
		}
	}
	return result, nil
}
