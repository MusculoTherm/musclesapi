package engine

import (
	"fmt"
	"github.com/MusculoTherm/musclesapi/models"
	"github.com/MusculoTherm/musclesapi/statistics"
	"sort"
	"strconv"
)

var workoutOnePoints []string = []string{"L0", "L1", "L2", "L3", "L4", "L5", "R0", "R1", "R2", "R3", "R4", "R5"}

func WorkoutOne(req models.WorkoutRequest) (models.WorkoutResponse, error) {
	resp := models.WorkoutResponse{}
	resp.TimeSpentS = req.TimeSpentS
	resp.PostPoints = calculateWorkoutOneImage(req.PostImage)
	resp.PrePoints = calculateWorkoutOneImage(req.PreImage)
	resp.DeltaPoints = models.FindDeltasForSetOfPoints(resp.PrePoints, resp.PostPoints)
	populateWorkoutOneBody(&resp)
	populateWorkoutOneTitle(&resp)
	resp.PreImage = req.PreImage.ImageURL
	resp.PostImage = req.PostImage.ImageURL
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
	var sumIncreased int64 = 0
	for _, p := range req.DeltaPoints {
		sumIncreased += p.MedianTemp
	}
	meanIncrease := statistics.Round(float64(sumIncreased) / float64(len(req.DeltaPoints)))
	sumMins := req.PrePoints[0].MinTemp
	for _, p := range req.DeltaPoints {
		sumMins += p.MinTemp
	}
	meanMins := statistics.Round(float64(sumMins) / float64(len(req.PrePoints)))
	sumMaxs := req.PrePoints[0].MaxTemp
	for _, p := range req.PostPoints {
		sumMins += p.MaxTemp
	}
	meanMaxs := statistics.Round(float64(sumMaxs) / float64(len(req.PostPoints)))
	req.Body += "Woah! What a great work out! The overall temperature of your targeted muscles increased by "
	req.Body += fmt.Sprintf("%d˚C, from %d˚ to %d˚. ", meanIncrease-273, meanMins-273, meanMaxs-273)
	lMedians := make([]int64, 6)
	rMedians := make([]int64, 6)
	for _, p := range req.PostPoints {
		alpha := p.Name[0:1]
		num, _ := strconv.Atoi(p.Name[1:2])
		if alpha == "R" {
			rMedians[num] = p.MedianTemp
		} else {
			lMedians[num] = p.MedianTemp
		}
	}
	diffHamstrings := float64((lMedians[0]+lMedians[1]))/2.0 - float64((rMedians[0]+rMedians[1]))/2.0
	if diffHamstrings > 0 {
		req.Body += fmt.Sprintf("In addition, we noticed that you worked your left hamstrings a bit harder than your right hamstrings. ")
	} else {
		req.Body += fmt.Sprintf("In addition, we noticed that you worked your right hamstrings a bit harder than your left hamstrings. ")
	}
	diffCalves := float64((lMedians[2]+lMedians[3]+lMedians[4]))/2.0 - float64((rMedians[2]+rMedians[3]+rMedians[4]))/2.0
	if diffCalves > 0 {
		req.Body += fmt.Sprintf("You also worked your left calves a bit harder than your right calves. ")
	} else {
		req.Body += fmt.Sprintf("You also worked your left calves a bit harder than your right calves. ")
	}
	diffAchilles := lMedians[5] - rMedians[5]
	if diffAchilles > -3 && diffAchilles < 3 {
		req.Body += fmt.Sprintf("On the other hand, your achilles are of nearly the same temperature. ")
	}
	if (float64((lMedians[0]+lMedians[1]))/2.0+float64((rMedians[0]+rMedians[1]))/2.0)/2.0 > (float64((lMedians[2]+lMedians[3]+lMedians[4]))/2.0+float64((rMedians[2]+rMedians[3]+rMedians[4]))/2.0)/2.0 {
		req.Body += fmt.Sprintf("During your work out, you warmed up your calves more than your hamstrings. Good job, if this was your goal, if not, work on targeting your hamstrings with some hamstring curls, or Romanian deadlift. ")
	} else {
		req.Body += fmt.Sprintf("During your work out, you warmed up your hamstrings more than your calves. Good job, if this was your goal, if not, work on targeting your calves more with some calf raises. ")
	}
	if req.TimeSpentS < 60 {
		req.Body += fmt.Sprintf("Next time, spend more time working out! You only worked out for 60 seconds, try for a few minutes next time! ")
	} else if req.TimeSpentS < 150 {
		req.Body += fmt.Sprintf("For spending less than two and a half minutes working out, this wasn't bad, next time try for a few minutes! ")
	} else {
		req.Body += fmt.Sprintf("Amazing job, you spent over two and a half minutes working out. Listen to these tips, and keep pushin' it! ")
	}
}
