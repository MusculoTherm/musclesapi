package statistics

import (
	"errors"
)

func Sum(input []float64) (sum float64, err error) {
	if len(input) == 0 {
		return 0, errors.New("Input must not be empty")
	}

	for _, n := range input {
		sum += n
	}

	return sum, nil
}
