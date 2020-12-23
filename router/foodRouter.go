package router

import (
	"github.com/gorilla/mux"
	FoodController "newapi.com/m/contollers"
)

//FoodRoutes function
func FoodRoutes(incomingRoutes *mux.Router) {

	// myRouter := mux.NewRouter().NewRoute().Subrouter().StrictSlash(true)

	incomingRoutes.HandleFunc("/foods", FoodController.GetFoods).Methods("GET")
	incomingRoutes.HandleFunc("/foods/{id}", FoodController.GetFood).Methods("GET")
	incomingRoutes.HandleFunc("/foods/{id}", FoodController.UpdateFood).Methods("PATCH")
	incomingRoutes.HandleFunc("/foods", FoodController.CreateFood).Methods("POST")

}
