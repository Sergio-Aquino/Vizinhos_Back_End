package Handler

import (
	"Vizinhos_Back_End/Entity"
	"Vizinhos_Back_End/Response"
	"encoding/json"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func GetSellerDataHandler(w http.ResponseWriter, r *http.Request) {
	// Database connection string
	dsn := "user=postgres password=postgres dbname=postgres port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	// Get seller ID from query parameters
	sellerID := strings.TrimPrefix(r.URL.Path, "/seller/")

	if sellerID == "" {
		http.Error(w, "Invalid Seller ID", http.StatusBadRequest)
		return
	}

	store := getSellerStore(sellerID, db)
	products := getSellerProducts(strconv.Itoa(store.StoreID), db)

	response := Response.SellerDataHandlerResponse{
		StoreAddress: store,
		Products:     products,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func getSellerStore(sellerID string, db *gorm.DB) Entity.StoreOrAddress {
	var store Entity.StoreOrAddress
	db.Joins("JOIN usuario ON usuario.fk_id_loja = loja_endereco.id_loja AND usuario.fk_id_endereco = loja_endereco.id_endereco").
		Where("usuario.cpf = ?", sellerID).Find(&store)

	store.CEP = strings.TrimSpace(store.CEP)

	return store
}

func getSellerProducts(storeID string, db *gorm.DB) []Entity.Product {
	var products []Entity.Product
	db.Where("fk_loja_endereco_id_loja = ?", storeID).Find(&products)
	return products
}
