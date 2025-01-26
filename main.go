package main

import (
	"Vizinhos_Back_End/Handler"
	"log"
)
import "net/http"

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

func main() {
	http.HandleFunc("/customer/", Handler.GetCustomerDataHandler)
	http.HandleFunc("/seller/", Handler.GetSellerDataHandler)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
