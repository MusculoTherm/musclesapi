package models

type ImagePointResponse struct {
	Name       string `json:"name"`
	X          int64  `json:"x"`
	Y          int64  `json:"y"`
	Radius     int64  `json:"r"`
	MaxTemp    int64  `json:"maxTempK"`
	Q1Temp     int64  `json:"q1TempK"`
	Q3Temp     int64  `json:"q3TempK"`
	MinTemp    int64  `json:"minTempK"`
	MeanTemp   int64  `json:"meanTempK"`
	MedianTemp int64  `json:"medianTempK"`
	IQR        int64  `json:"IQRK"`
}
