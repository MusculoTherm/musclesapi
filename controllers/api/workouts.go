package api

import (
	"github.com/MusculoTherm/musclesapi/models"
	"net/http"
)

func V0_API_Create_Workout(w http.ResponseWriter, r *http.Request) {
	resp := models.WorkoutResponse{}
	success := models.Response{Success: true, Data: resp, Message: "Workout Success"}
	JSONResponse(w, success, 200)
}
