package controllers

import (
	"github.com/MusculoTherm/musclesapi/controllers/api"
	"github.com/gorilla/mux"
	"net/http"
)

func CreateRouter() http.Handler {
	router := mux.NewRouter()
	apiV0Router := router.PathPrefix("/v0").Subrouter()
	apiV0Router = apiV0Router.StrictSlash(true)
	apiV0Router.HandleFunc("/", Use(api.V0_API)).Methods("GET")
	apiV0Router.HandleFunc("/workouts", Use(api.V0_API_Create_Workout)).Methods("POST")
	apiV0Router.HandleFunc("/uploads", Use(api.V0_API_Upload_Artwork)).Methods("POST")
	router.HandleFunc("/workouts", ServeWorkout).Methods("GET")
	uploadsFS := http.StripPrefix("/uploads/", http.FileServer(http.Dir("./uploads/")))
	router.PathPrefix("/uploads/").Handler(uploadsFS)
	router.HandleFunc("/", ServeHome)
	return router
}

// `Use` allows us to stack middleware to process the request
// Example taken from https://github.com/gorilla/mux/pull/36#issuecomment-25849172
func Use(handler http.HandlerFunc, mid ...func(http.Handler) http.HandlerFunc) http.HandlerFunc {
	for _, m := range mid {
		handler = m(handler)
	}
	return handler
}
