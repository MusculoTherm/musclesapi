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

func I64ToI(in []int64) []int {
	out := make([]int, len(in))
	for ind, i := range in {
		out[ind] = int(i)
	}
	return out
}

func IToI64(in []int) []int64 {
	out := make([]int64, len(in))
	for ind, i := range in {
		out[ind] = int64(i)
	}
	return out
}

func I64ToF64(in []int64) []float64 {
	out := make([]float64, len(in))
	for ind, i := range in {
		out[ind] = float64(i)
	}
	return out
}

func F64ToI64(in []float64) []int64 {
	out := make([]int64, len(in))
	for ind, i := range in {
		out[ind] = Round(i)
	}
	return out
}
