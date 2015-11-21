package statistics
import "fmt"

func GetElementsWithinRadius(xCenter int, yCenter int, radius int, mat [][]int64) []int64 {
	out := make([]int64, 0, radius*radius)
	for x := xCenter - radius; x <= xCenter; x++ {
		for y := yCenter - radius; y <= yCenter; y++ {
			if (x-xCenter)*(x-xCenter)+(y-yCenter)*(y-yCenter) < radius*radius {
				if y > -1 && x > -1 && y < len(mat) && x < len(mat[y]) {
					if mat[y][x] <= 305 {
						out = append(out, mat[y][x])
					}
				}
			}
		}
	}
	fmt.Println(len(out))
	return out
}
