package models

type WorkoutRequest struct {
	WorkoutId  int64               `json:"workoutId"`
	TimeSpentS int64               `json:"timeSpentSeconds"`
	PreImage   ImageDetailsRequest `json:"pre"`
	PostImage  ImageDetailsRequest `json:"post"`
}

func (req *WorkoutRequest) Parse() {
	h := int(req.PostImage.Height)
	w := int(req.PostImage.Width)
	tempMat := make([][]int64, h)
	for i := 0; i < h; i++ {
		tempMat[i] = make([]int64, w)
	}
	pos := 0
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			tempMat[i][j] = req.PostImage.Temps[pos]
			pos++
		}
	}
	req.PostImage.TempsMat = tempMat
}
