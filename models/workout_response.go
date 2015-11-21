package models

import "time"

type WorkoutResponse struct {
	Id          int64                `json:"id"`
	Endpoint    string               `json:"endpoint"`
	ApiEndpoint string               `json:"apiEndpoint"`
	PreImage    string               `json:"preImageUrl"`
	PostImage   string               `json:"postImageUrl"`
	PrePoints   []ImagePointResponse `json:"prePoints"`
	PostPoints  []ImagePointResponse `json:"postPoints"`
	DeltaPoints []ImagePointDeltas   `json:"deltaPoints"`
	Title       string               `json:"title"`
	Body        string               `json:"body"`
	TimeSpentS  int64                `json:"timeSpentSeconds"`
	CreatedAt   time.Time            `json:"createdAt"`
	UpdatedAt   time.Time            `json:"updatedAt"`
}

func (workout *WorkoutResponse) FindById() error {
	raw := Workout{Id: workout.Id}
	err := raw.FindById()
	if err != nil {
		return err
	}
	raw.PopulateResponse(workout)
	return nil
}

func (workout *WorkoutResponse) Create() error {
	raw := Workout{}
	raw.Populate(workout)
	err := raw.Create()
	if err != nil {
		return err
	}
	raw.PopulateResponse(workout)
	return nil
}
