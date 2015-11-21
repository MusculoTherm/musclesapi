package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/MusculoTherm/musclesapi/models"
)

// API (/api/v0) provides access to a simple screen to ensure the server is up
func V0_API(w http.ResponseWriter, r *http.Request) {
	hello := models.Response{Success: true, Debug: "Server is healthy", Message: "Server is healthy"}
	JSONResponse(w, hello, 200)
}

// JSONResponse attempts to set the status code, c, and marshal the given interface, d, into a response that
// is written to the given ResponseWriter.
func JSONResponse(w http.ResponseWriter, d interface{}, c int) {
	//dj, err := json.MarshalIndent(d, "", "  ")
	dj, err := json.Marshal(d)
	if err != nil {
		http.Error(w, "Error creating JSON response", http.StatusInternalServerError)
		fmt.Println(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(c)
	fmt.Fprintf(w, "%s", dj)
}
