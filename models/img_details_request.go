package models

type ImageDetailsRequest struct {
	ImageURL  string `json:"imageURL"`
	Temps     []int64 `json:"tempsK"`
	Dimensions Dimensions `json:"dimensions"`
}