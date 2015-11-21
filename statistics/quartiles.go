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
	ind1 := int(float64(len(orig)) * 0.75)
	ind2 := int(float64(len(orig))*0.75 + 1.0)
	m, _ := Mean(I64ToF64(orig[ind1:ind2]))
	return Round(m)
}
