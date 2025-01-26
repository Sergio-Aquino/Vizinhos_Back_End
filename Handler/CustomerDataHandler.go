package Handler

import (
	"Vizinhos_Back_End/Entity"
	"Vizinhos_Back_End/Response"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func GetCustomerDataHandler(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimPrefix(r.URL.Path, "/customer/")
	customerID, err := strconv.Atoi(path)
	if err != nil {
		http.Error(w, "Invalid Customer ID", http.StatusBadRequest)
		return
	}

	orders := getCustomerOrders(customerID)
	addresses := getCustomerAddresses(customerID)

	response := Response.CustomerDataHandlerResponse{
		Orders:    orders,
		Addresses: addresses,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func getCustomerOrders(customerID int) []Entity.Order {
	return []Entity.Order{
		{
			OrderID: 1,
			User: Entity.User{
				CPF:      "47713289730",
				UserType: 0,
				StoreOrAddress: []Entity.StoreOrAddress{{
					StoreID:    1,
					AddressID:  1,
					CEP:        "12345678",
					Street:     "Street 1",
					Number:     123,
					Complement: "Complement 1",
					StoreName:  "Store 1",
				}},
				PhoneNumber:  "123456789",
				Email:        "client1@gmail.com",
				RegisterDate: time.Now(),
			},
			Batch: Entity.Batch{
				BatchID: 1,
				Product: Entity.Product{
					ProductID: 1,
					Store: Entity.StoreOrAddress{
						StoreID:    1,
						AddressID:  1,
						CEP:        "12345678",
						Street:     "Street 1",
						Number:     123,
						Complement: "Complement 1",
						StoreName:  "Store 1",
					},
					Category: Entity.Category{
						CategoryID:  1,
						Description: "Category 1",
					},
					DaysToExp:          30,
					SellPrice:          100.0,
					ManufacturingPrice: 50.0,
					Size:               1.0,
				},
				ManufacturingDate: time.Now(),
				Discount:          10.0,
				Quantity:          100,
			},
			Price:      100.0,
			Quantity:   2,
			Date:       time.Now(),
			Status:     "Delivered",
			LastUpdate: time.Now(),
		},
	}
}

func getCustomerAddresses(customerID int) []Entity.StoreOrAddress {
	return []Entity.StoreOrAddress{
		{
			StoreID:    1,
			AddressID:  1,
			CEP:        "12345678",
			Street:     "Street 1",
			Number:     123,
			Complement: "Complement 1",
		},
		{
			StoreID:    2,
			AddressID:  1,
			CEP:        "68234777",
			Street:     "Street 2",
			Number:     124,
			Complement: "Complement 2",
		},
	}
}
