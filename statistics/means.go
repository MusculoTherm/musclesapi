package statistics

import (
	"errors"
	"fmt"
)

func Mean(input []float64) (float64, error) {
	if len(input) == 0 {
		return 0, errors.New("Input must not be empty")
	}

	sum, _ := Sum(input)

	return sum / float64(len(input)), nil
}

func MatrixMeans(input [][]float64) ([]float64, error) {
	result := make([]float64, len(input[0]))
	for y, _ := range input[0] {
		var err error
		result[y], err = Mean(getColumn(input, y))
		if err != nil {
			return result, errors.New(fmt.Sprintf("Error calculating means for matrix: ", err))
		}
	}
	return result, nil
}
