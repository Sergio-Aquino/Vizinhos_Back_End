package main

import (
	"Vizinhos_Back_End/Handler"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"net/http"
)

func main() {
	// Database connection string
	dsn := "user=postgres password=postgres dbname=postgres port=5432 sslmode=disable"
	_, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	// Set up HTTP handlers
	http.HandleFunc("/customer/", Handler.GetCustomerDataHandler)
	http.HandleFunc("/seller/", Handler.GetSellerDataHandler)

	// Start the HTTP server
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
