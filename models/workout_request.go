package models

type WorkoutRequest struct {
	WorkoutId  int64               `json:"workoutId"`
	TimeSpentS int64               `json:"timeSpentSeconds"`
	PreImage   ImageDetailsRequest `json:"pre"`
	PostImage  ImageDetailsRequest `json:"post"`
}
