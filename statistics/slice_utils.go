package statistics

import ()

func copyslice(input []float64) []float64 {
	s := make([]float64, len(input))
	copy(s, input)
	return s
}

func getColumn(input [][]float64, index int) []float64 {
	output := make([]float64, len(input))
	for y, row := range input {
		output[y] = row[index]
	}
	return output
}
