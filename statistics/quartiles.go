package statistics

func GetQ1FromSortedInts(orig []int64) int64 {
	q := 0.0
	input := I64ToF64(orig)
	l := len(input)
	if l%4 == 0 {
		q, _ = Mean(input[l/4-1 : l/4+1])
	} else {
		q = float64(input[l/4])
	}
	return Round(q)
}

func GetQ3FromSortedInts(orig []int64) int64 {
	//TODO
	return GetQ1FromSortedInts(orig)
}