package models

type ImagePointDeltas struct {
	Name       string `json:"name"`
	MaxTemp    int64  `json:"makeTempK"`
	Q1Temp     int64  `json:"q1TempK"`
	Q3Temp     int64  `json:"q3TempK"`
	MinTemp    int64  `json:"minTempK"`
	MeanTemp   int64  `json:"meanTempK"`
	MedianTemp int64  `json:"medianTempK`
	IQR        int64  `json:"IQRK"`
}

func FindDeltasForSetOfPoints(pre []ImagePointResponse, post []ImagePointResponse) []ImagePointDeltas {
	resp := make([]ImagePointDeltas, 0, len(post))
	for i := 0; i < len(pre); i++ {
		for j := 0; j < len(post); j++ {
			if pre[i].Name == post[j].Name {
				resp = append(resp, FindDeltasForPoint(pre[i], post[j]))
			}
		}
	}
	return resp
}

func FindDeltasForPoint(pre ImagePointResponse, post ImagePointResponse) ImagePointDeltas {
	return ImagePointDeltas{
		Name:       pre.Name,
		MaxTemp:    pre.MaxTemp - post.MaxTemp,
		MinTemp:    pre.MinTemp - post.MinTemp,
		Q1Temp:     pre.Q1Temp - post.Q1Temp,
		Q3Temp:     pre.Q3Temp - post.Q3Temp,
		MeanTemp:   pre.MeanTemp - post.MeanTemp,
		MedianTemp: pre.MedianTemp - post.MedianTemp,
		IQR:        pre.IQR - post.IQR,
	}
}
