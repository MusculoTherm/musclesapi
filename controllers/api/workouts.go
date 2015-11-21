package api

import (
	"encoding/json"
	"fmt"
	"github.com/MusculoTherm/musclesapi/engine"
	"github.com/MusculoTherm/musclesapi/models"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func V0_API_Create_Workout(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var wReq models.WorkoutRequest
	err := decoder.Decode(&wReq)
	if err != nil {
		failFormat := models.Response{Success: false, Debug: "Check JSON Formatting", Message: "Workout Failure"}
		JSONResponse(w, failFormat, 400)
		fmt.Println(err)
		return
	}
	wReq.Parse()
	resp, err := engine.Execute(wReq)
	if err != nil {
		failFormat := models.Response{Success: false, Debug: "Internal Server Error", Message: "Workout Failure"}
		JSONResponse(w, failFormat, 500)
		fmt.Println(err)
		return
	}
	err = resp.Create()
	if err != nil {
		failFormat := models.Response{Success: false, Debug: "Internal Server Error [DATABASE]", Message: "Workout Failure"}
		JSONResponse(w, failFormat, 500)
		fmt.Println(err)
		return
	}
	success := models.Response{Success: true, Data: resp, Message: "Workout Success"}
	JSONResponse(w, success, 200)
}

func V0_API_Get_Workout(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	wIdStr := vars["workoutId"]
	wId, err := strconv.ParseInt(wIdStr, 10, 64)
	if err != nil {
		failParams := models.Response{Success: false, Debug: fmt.Sprintf("Invalid workout id, %s", wIdStr), Message: "Failed Retrieving the Workout"}
		JSONResponse(w, failParams, 400)
		return
	}
	pl := models.WorkoutResponse{Id: wId}
	err = pl.FindById()
	if err != nil {
		failParams := models.Response{Success: false, Debug: "Internal server error. If this issue persists, please submit a bug report to {{serviceName}}", Message: "Failed Retrieving the Workout"}
		JSONResponse(w, failParams, 500)
		return
	}
	successResp := models.Response{Success: true, Data: pl, Message: "Successfully Retrieved the Workout"}
	JSONResponse(w, successResp, 200)
}
