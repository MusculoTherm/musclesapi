package models

type WorkoutResponse struct {
	Endpoint   string               `json:"endpoint"`
	PrePoints  []ImagePointResponse `json:"prePoints"`
	PostPoints []ImagePointResponse `json:"postPoints"`
}
