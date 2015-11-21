package models

type ImageDetailsRequest struct {
	ImageURL string              `json:"imageURL"`
	Temps    []int64             `json:"tempsK"`
	TempsMat [][]int64           `json:"-"`
	Width    int64               `json:"width"`
	Height   int64               `json:"height"`
	Points   []ImagePointRequest `json:"points"`
}
