package models

type WorkoutRequest struct {
	WorkoutId  int64               `json:"workoutId"`
	TimeSpentS int64               `json:"timeSpentSeconds"`
	PreImage   ImageDetailsRequest `json:"pre"`
	PostImage  ImageDetailsRequest `json:"post"`
}

func (req *WorkoutRequest) Parse() {
	req.PostImage.TempsMat = parseImageToMat(req.PostImage)
	req.PreImage.TempsMat = parseImageToMat(req.PreImage)
}

func parseImageToMat(r ImageDetailsRequest) [][]int64 {
	h := int(r.Height)
	w := int(r.Width)
	tempMat := make([][]int64, h)
	for i := 0; i < h; i++ {
		tempMat[i] = make([]int64, w)
	}
	pos := 0
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			tempMat[i][j] = r.Temps[pos]
			pos++
		}
	}
	return tempMat
}
