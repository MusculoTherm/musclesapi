package controllers

import (
	"fmt"
	"github.com/MusculoTherm/musclesapi/controllers/api"
	"github.com/MusculoTherm/musclesapi/models"
	"github.com/gorilla/mux"
	"html/template"
	"net/http"
	"strconv"
)

var workoutTempl = template.Must(template.ParseFiles("templates/workout.html"))

func ServeWorkout(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	wIdStr := vars["workoutId"]
	wId, err := strconv.ParseInt(wIdStr, 10, 64)
	if err != nil {
		failParams := models.Response{Success: false, Debug: fmt.Sprintf("Invalid workout id, %s", wIdStr), Message: "Failed Retrieving the Workout"}
		api.JSONResponse(w, failParams, 400)
		return
	}
	pl := models.WorkoutResponse{Id: wId}
	err = pl.FindById()
	if err != nil {
		failParams := models.Response{Success: false, Debug: "Internal server error. If this issue persists, please submit a bug report to {{serviceName}}", Message: "Failed Retrieving the Workout"}
		api.JSONResponse(w, failParams, 500)
		return
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	workoutTempl.Execute(w, pl)
}
