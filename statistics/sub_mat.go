package statistics

func GetElementsWithinRadius(xCenter int, yCenter int, radius int, mat [][]int64) []int64 {
	out := make([]int64, 0, radius*radius)
	for x := xCenter - radius; x <= xCenter; x++ {
		for y := yCenter - radius; y <= yCenter; y++ {
			if (x-xCenter)*(x-xCenter)+(y-yCenter)*(y-yCenter) < radius*radius {
				if y > -1 && x > -1 && y < len(mat) && x < len(mat[y]) {
					out = append(out, mat[y][x])
				}
			}
		}
	}
	return out
}
