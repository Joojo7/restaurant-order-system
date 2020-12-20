package router

import (
	"github.com/gorilla/mux"
	MenuController "newapi.com/m/contollers"
)

//Routes function
func Routes(incomingRoutes *mux.Router) {

	// myRouter := mux.NewRouter().NewRoute().Subrouter().StrictSlash(true)

	incomingRoutes.HandleFunc("/menus", MenuController.GetMenus).Methods("GET")
	incomingRoutes.HandleFunc("/menus/{id}", MenuController.GetMenu).Methods("GET")
	// incomingRoutes.HandleFunc("/menus/{id}", MenuController.UpdateMenu).Methods("PATCH")
	incomingRoutes.HandleFunc("/menus", MenuController.CreateMenu).Methods("POST")

}
