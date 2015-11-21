package engine

import (
	"encoding/json"
	"fmt"
	"github.com/MusculoTherm/musclesapi/models"
	"github.com/MusculoTherm/musclesapi/statistics"
	"sort"
)

var workoutOnePoints []string = []string{"L0", "L1", "L2", "L3", "L4", "L5", "R0", "R1", "R2", "R3", "R4", "R5"}

func WorkoutOne(req models.WorkoutRequest) (models.WorkoutResponse, error) {
	resp := models.WorkoutResponse{}
	resp.PostPoints = calculateWorkoutOneImage(req.PostImage)
	resp.PrePoints = calculateWorkoutOneImage(req.PreImage)
	resp.DeltaPoints = models.FindDeltasForSetOfPoints(resp.PrePoints, resp.PostPoints)
	populateWorkoutOneBody(&resp)
	populateWorkoutOneTitle(&resp)
	return resp, nil
}

func calculateWorkoutOneImage(req models.ImageDetailsRequest) []models.ImagePointResponse {
	resp := models.MirrorImagePointRequestsToResponses(req.Points)
	for ind, p := range resp {
		points64Unsorted := statistics.GetElementsWithinRadius(int(p.X), int(p.Y), int(p.Radius), req.TempsMat)
		points32 := statistics.I64ToI(points64Unsorted)
		sort.Ints(points32)
		points := statistics.IToI64(points32)
		resp[ind].MaxTemp = maxTempForPoint(points)
		resp[ind].MinTemp = minTempForPoint(points)
		resp[ind].Q1Temp = q1TempForPoint(points)
		resp[ind].Q3Temp = q3TempForPoint(points)
		resp[ind].MeanTemp = meanTempForPoint(points)
		resp[ind].MedianTemp = medianTempForPoint(points)
		resp[ind].IQR = iqrTempForPoint(points)
	}
	return resp
}

func populateWorkoutOneTitle(req *models.WorkoutResponse) {
	req.Title = "Good Job, Mate!"
}

func populateWorkoutOneBody(req *models.WorkoutResponse) {
	b, _ := json.Marshal(req.PrePoints)
	req.Body += fmt.Sprintf("Pre: %s ", string(b))
	b, _ = json.Marshal(req.PostPoints)
	req.Body += fmt.Sprintf("||||Post: %s ", string(b))
	b, _ = json.Marshal(req.PostPoints)
	req.Body += fmt.Sprintf("||||Deltas: %s ", string(b))
}
