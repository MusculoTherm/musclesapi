package models

import (
	"encoding/json"
	"fmt"
	"time"
)

type Workout struct {
	Id        int64     `gorm:"primary_key"sql:"AUTO_INCREMENT"`
	Value     string    `sql:"type:text"`
	CreatedAt time.Time `sql:"type:datetime"`
	UpdatedAt time.Time `sql:"type:datetime"`
}

func (workout *Workout) FindById() error {
	return db.Where(`id = ?`, workout.Id).First(&workout).Error
}

func (workout *Workout) Create() error {
	return db.Create(&workout).Error
}

func (workout *Workout) Populate(data *WorkoutResponse) {
	b, err := json.Marshal(&data)
	if err != nil {
		panic(err)
	}
	workout.Value = string(b)
}

func (workout *Workout) PopulateResponse(resp *WorkoutResponse) {
	err := json.Unmarshal([]byte(workout.Value), &resp)
	if err != nil {
		panic(err)
	}
	resp.ApiEndpoint = fmt.Sprintf("%sv0/workouts/%d", GlobalConfig.FullyQualifiedHost, workout.Id)
	resp.Endpoint = fmt.Sprintf("%sworkouts/%d", GlobalConfig.FullyQualifiedHost, workout.Id)
	resp.Id = workout.Id
	resp.CreatedAt = workout.CreatedAt
	resp.UpdatedAt = workout.UpdatedAt
}
