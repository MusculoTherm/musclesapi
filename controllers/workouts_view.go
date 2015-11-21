package controllers

import (
	"net/http"
	"html/template"
	"github.com/gorilla/mux"
	"strconv"
"github.com/MusculoTherm/musclesapi/models"
	"fmt"
"github.com/MusculoTherm/musclesapi/controllers/api"
)

var workoutTempl = template.Must(template.ParseFiles("templates/workout.html"))

func ServeWorkout(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "Not found", 404)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", 405)
		return
	}
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
	homeTempl.Execute(w, pl)
}
