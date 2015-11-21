package api

import (
	"encoding/json"
	"fmt"
	"github.com/MusculoTherm/musclesapi/models"
	"net/http"
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
	resp := models.WorkoutResponse{}
	success := models.Response{Success: true, Data: resp, Message: "Workout Success"}
	JSONResponse(w, success, 200)
}
