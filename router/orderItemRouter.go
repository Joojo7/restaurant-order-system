package router

import (
	"github.com/gorilla/mux"
	OrderItemController "newapi.com/m/contollers"
)

//OrderItemRoutes function
func OrderItemRoutes(incomingRoutes *mux.Router) {

	// myRouter := mux.NewRouter().NewRoute().Subrouter().StrictSlash(true)

	incomingRoutes.HandleFunc("/orderItems", OrderItemController.GetOrderItems).Methods("GET")
	incomingRoutes.HandleFunc("/orderItems/{id}", OrderItemController.GetOrderItem).Methods("GET")
	incomingRoutes.HandleFunc("/orderItems/{id}", OrderItemController.UpdateOrderItem).Methods("PATCH")
	incomingRoutes.HandleFunc("/orderItems", OrderItemController.CreateOrderItem).Methods("POST")

}
