package router

import (
	"github.com/gorilla/mux"
	CoasterController "newapi.com/m/contollers"
	NoteController "newapi.com/m/contollers"
)

//Routes function
func Routes(incomingRoutes *mux.Router) {

	coasterHandlers := CoasterController.NewCoasterHandlers()
	// myRouter := mux.NewRouter().NewRoute().Subrouter().StrictSlash(true)
	incomingRoutes.HandleFunc("/coasters", coasterHandlers.Get).Methods("GET")
	incomingRoutes.HandleFunc("/notes", NoteController.GetNotes).Methods("GET")
	incomingRoutes.HandleFunc("/notes/{id}", NoteController.GetNote).Methods("GET")
	incomingRoutes.HandleFunc("/notes/{id}", NoteController.UpdateNote).Methods("PATCH")
	incomingRoutes.HandleFunc("/notes", NoteController.CreateNote).Methods("POST")
	incomingRoutes.HandleFunc("/coasters", coasterHandlers.Post).Methods("POST")
}
