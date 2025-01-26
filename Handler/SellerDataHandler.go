package Handler

import (
	"Vizinhos_Back_End/Entity"
	"Vizinhos_Back_End/Response"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

func GetSellerDataHandler(w http.ResponseWriter, r *http.Request) {
	sellerID, err := strconv.Atoi(strings.TrimPrefix(r.URL.Path, "/seller/"))

	if err != nil {
		http.Error(w, "Invalid Seller ID", http.StatusBadRequest)
		return
	}

	store := getSellerStore(sellerID)
	products := getSellerProducts(store.StoreID)

	response := Response.SellerDataHandlerResponse{
		StoreAddress: store,
		Products:     products,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

}

func getSellerStore(sellerID int) Entity.StoreOrAddress {
	return Entity.StoreOrAddress{
		StoreID:    1,
		AddressID:  1,
		CEP:        "12345678",
		Street:     "Street 1",
		Number:     123,
		Complement: "Complement 1",
		StoreName:  "Store 1",
	}
}

func getSellerProducts(storeID int) []Entity.Product {
	return []Entity.Product{
		{
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
	}
}
