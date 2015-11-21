package models

type WorkoutResponse struct {
	Endpoint int64 `json:"endpoint"`
	PrePoints []ImagePointResponse `json:"prePoints"`
	PostPoints []ImagePointResponse `json:"postPoints"`
}