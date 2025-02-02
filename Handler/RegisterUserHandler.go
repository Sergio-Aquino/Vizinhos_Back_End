package Handler

import (
	"Vizinhos_Back_End/Entity"
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"net/http"
	"time"
)

func RegisterUserHandler(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	dsn := "user=postgres password=postgres dbname=postgres port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	var user Entity.User
	err = json.Unmarshal([]byte(req.Body), &user)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusBadRequest,
			Body:       "Invalid request payload",
		}, nil
	}

	if user.UserType != 0 && user.UserType != 1 {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusBadRequest,
			Body:       "Invalid user type",
		}, nil
	}

	user.RegisterDate = time.Now()

	if user.UserType == 1 {
		user.StoreID = 0

		var lastAddressID int
		db.Table("loja_endereco").Select("MAX(id_endereco)").Scan(&lastAddressID)
		user.StoreOrAddress.AddressID = lastAddressID + 1

		err = db.Create(&user.StoreOrAddress).Error
		if err != nil {
			return events.APIGatewayProxyResponse{
				StatusCode: http.StatusInternalServerError,
				Body:       "Failed to register user address",
			}, nil
		}

		user.AddressID = user.StoreOrAddress.AddressID

	} else {
		var lastStoreID int
		db.Table("loja_endereco").Select("MAX(id_loja)").Scan(&lastStoreID)
		user.StoreOrAddress.StoreID = lastStoreID + 1

		var lastAddressID int
		db.Table("loja_endereco").Select("MAX(id_endereco)").Scan(&lastAddressID)
		user.StoreOrAddress.AddressID = lastAddressID + 1

		err = db.Create(&user.StoreOrAddress).Error
		if err != nil {
			return events.APIGatewayProxyResponse{
				StatusCode: http.StatusInternalServerError,
				Body:       "Failed to register seller store",
			}, nil
		}

		user.StoreID = user.StoreOrAddress.StoreID
		user.AddressID = user.StoreOrAddress.AddressID
	}

	err = db.Create(&user).Error
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       "Failed to register user",
		}, nil
	}

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusCreated,
		Body:       "User registered successfully",
	}, nil
}
