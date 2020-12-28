package router

import (
	"github.com/gorilla/mux"
	OrderController "newapi.com/m/contollers"
)

//OrderRoutes function
func OrderRoutes(incomingRoutes *mux.Router) {

	// myRouter := mux.NewRouter().NewRoute().Subrouter().StrictSlash(true)

	incomingRoutes.HandleFunc("/orders", OrderController.GetOrders).Methods("GET")
	incomingRoutes.HandleFunc("/orders/{id}", OrderController.GetOrder).Methods("GET")
	incomingRoutes.HandleFunc("/orders/{id}", OrderController.UpdateOrder).Methods("PATCH")
	incomingRoutes.HandleFunc("/orders", OrderController.CreateOrder).Methods("POST")

}
