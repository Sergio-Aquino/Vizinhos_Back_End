package main

import (
	"Vizinhos_Back_End/Handler"
	"log"
	"net/http"
)

func main() {
	// Set up HTTP handlers
	http.HandleFunc("/customer/", Handler.GetCustomerDataHandler)
	http.HandleFunc("/seller/", Handler.GetSellerDataHandler)
	http.HandleFunc("/register/user/", Handler.RegisterUserHandler)

	// Start the HTTP server
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
