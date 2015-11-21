package models

type ImagePointRequest struct {
	Name   string `json:"name"`
	X      int64  `json:"x"`
	Y      int64  `json:"y"`
	Radius int64  `json:"r"`
}
