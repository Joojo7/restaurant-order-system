package router

import (
	"github.com/gorilla/mux"
	InvoiceController "newapi.com/m/contollers"
)

//InvoiceRoutes function
func InvoiceRoutes(incomingRoutes *mux.Router) {

	// myRouter := mux.NewRouter().NewRoute().Subrouter().StrictSlash(true)

	incomingRoutes.HandleFunc("/invoices", InvoiceController.GetInvoices).Methods("GET")
	incomingRoutes.HandleFunc("/invoices/{id}", InvoiceController.GetInvoice).Methods("GET")
	incomingRoutes.HandleFunc("/invoices/{id}", InvoiceController.UpdateInvoice).Methods("PATCH")
	incomingRoutes.HandleFunc("/invoices", InvoiceController.CreateInvoice).Methods("POST")

}
