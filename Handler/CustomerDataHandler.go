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
	"strings"
)

func GetCustomerDataHandler(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	dsn := "user=postgres password=postgres dbname=postgres port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	customerID := strings.TrimPrefix(req.Path, "/customer/")
	if customerID == "" {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusBadRequest,
			Body:       "Invalid customer ID",
		}, nil
	}

	orders := getCustomerOrders(customerID, db)
	address := getCustomerAddresses(customerID, db)

	response := Response.CustomerDataHandlerResponse{
		Orders:  orders,
		Address: address,
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

func getCustomerOrders(customerID string, db *gorm.DB) []Entity.Order {
	var orders []Entity.Order
	db.Where("fk_usuario_cpf = ?", strings.TrimSpace(customerID)).Find(&orders)

	for i := range orders {
		orders[i].UserCPF = strings.TrimSpace(orders[i].UserCPF)
	}

	return orders
}

func getCustomerAddresses(customerID string, db *gorm.DB) Entity.StoreOrAddress {
	var address Entity.StoreOrAddress
	db.Joins("JOIN usuario ON usuario.fk_id_loja = loja_endereco.id_loja AND usuario.fk_id_endereco = loja_endereco.id_endereco").
		Where("usuario.cpf = ?", customerID).Find(&address)
	address.CEP = strings.TrimSpace(address.CEP)

	return address
}
