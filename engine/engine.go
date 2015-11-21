package engine

import (
	"errors"
	"github.com/MusculoTherm/musclesapi/models"
)

func Execute(req models.WorkoutRequest) (models.WorkoutResponse, error) {
	switch req.WorkoutId {
	case 1:
		return WorkoutOne(req)
	default:
		return models.WorkoutResponse{}, errors.New("Unknown workout ID")
	}
}
