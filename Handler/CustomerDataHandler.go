package Handler

import (
	"Vizinhos_Back_End/Entity"
	"Vizinhos_Back_End/Response"
	"encoding/json"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"net/http"
	"strings"
)

func GetCustomerDataHandler(w http.ResponseWriter, r *http.Request) {
	// Database connection string
	dsn := "user=postgres password=postgres dbname=postgres port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	// Get customer ID from query parameters
	customerID := strings.TrimPrefix(r.URL.Path, "/customer/")
	if customerID == "" {
		http.Error(w, "Invalid customer ID", http.StatusBadRequest)
		return
	}

	orders := getCustomerOrders(customerID, db)
	addresses := getCustomerAddresses(customerID, db)

	response := Response.CustomerDataHandlerResponse{
		Orders:    orders,
		Addresses: addresses,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func getCustomerOrders(customerID string, db *gorm.DB) []Entity.Order {
	var orders []Entity.Order
	db.Where("fk_usuario_cpf = ?", strings.TrimSpace(customerID)).Find(&orders)

	for i := range orders {
		orders[i].UserCPF = strings.TrimSpace(orders[i].UserCPF)
	}

	return orders
}

func getCustomerAddresses(customerID string, db *gorm.DB) []Entity.StoreOrAddress {
	var addresses []Entity.StoreOrAddress
	db.Joins("JOIN usuario ON usuario.fk_id_loja = loja_endereco.id_loja AND usuario.fk_id_endereco = loja_endereco.id_endereco").
		Where("usuario.cpf = ?", customerID).Find(&addresses)
	return addresses
}
