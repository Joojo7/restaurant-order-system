package router

import (
	"github.com/gorilla/mux"
	TableController "newapi.com/m/contollers"
)

//TableRoutes function
func TableRoutes(incomingRoutes *mux.Router) {

	// myRouter := mux.NewRouter().NewRoute().Subrouter().StrictSlash(true)

	incomingRoutes.HandleFunc("/tables", TableController.GetTables).Methods("GET")
	incomingRoutes.HandleFunc("/tables/{id}", TableController.GetTable).Methods("GET")
	incomingRoutes.HandleFunc("/tables/{id}", TableController.UpdateTable).Methods("PATCH")
	incomingRoutes.HandleFunc("/tables", TableController.CreateTable).Methods("POST")

}
