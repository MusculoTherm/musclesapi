package engine

import "github.com/MusculoTherm/musclesapi/models"

var workoutOnePoints []string = []string{"L0", "L1", "L2", "L3", "L4", "L5", "R0", "R1", "R2", "R3", "R4", "R5"}

func WorkoutOne(req models.WorkoutRequest) (models.WorkoutResponse, error) {
	resp := models.WorkoutResponse{}
	resp.PostPoints = calculateWorkoutOneImage(req.PostImage)
	resp.PrePoints = calculateWorkoutOneImage(req.PreImage)
	resp.DeltaPoints = models.FindDeltasForSetOfPoints(resp.PrePoints, resp.PostPoints)
	return resp, nil
}

func calculateWorkoutOneImage(req models.ImageDetailsRequest) []models.ImagePointResponse {
	resp := models.MirrorImagePointRequestsToResponses(req.Points)
	for ind, p := range resp {
		resp[ind].MaxTemp = maxTempForPoint(p.X, p.Y, p.Radius, req.TempsMat)
		resp[ind].MinTemp = minTempForPoint(p.X, p.Y, p.Radius, req.TempsMat)
		resp[ind].Q1Temp = q1TempForPoint(p.X, p.Y, p.Radius, req.TempsMat)
		resp[ind].Q3Temp = q3TempForPoint(p.X, p.Y, p.Radius, req.TempsMat)
		resp[ind].MeanTemp = meanTempForPoint(p.X, p.Y, p.Radius, req.TempsMat)
		resp[ind].MedianTemp = medianTempForPoint(p.X, p.Y, p.Radius, req.TempsMat)
		resp[ind].IQR = iqrTempForPoint(p.X, p.Y, p.Radius, req.TempsMat)
	}
	return resp
}
