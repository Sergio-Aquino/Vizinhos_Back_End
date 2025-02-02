package Handler

import (
	"Vizinhos_Back_End/Entity"
	"Vizinhos_Back_End/Response"
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func GetSellerDataHandler(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	dsn := "user=postgres password=postgres dbname=postgres port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	sellerID := strings.TrimPrefix(req.Path, "/seller/")
	if sellerID == "" {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusBadRequest,
			Body:       "Invalid Seller ID",
		}, nil
	}

	store := getSellerStore(sellerID, db)
	products := getSellerProducts(strconv.Itoa(store.StoreID), db)

	response := Response.SellerDataHandlerResponse{
		StoreAddress: store,
		Products:     products,
	}

	responseBody, err := json.Marshal(response)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       "Failed to marshal response",
		}, nil
	}

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Body:       string(responseBody),
	}, nil
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
