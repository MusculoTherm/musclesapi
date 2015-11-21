package models

type ImageDetailsRequest struct {
	ImageURL string  `json:"imageURL"`
	Temps    []int64 `json:"tempsK"`
	Width    int64   `json:"width"`
	Height   int64   `json:"height"`
}
