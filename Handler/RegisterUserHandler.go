package Handler

import (
	"Vizinhos_Back_End/Entity"
	"encoding/json"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"net/http"
	"time"
)

func RegisterUserHandler(w http.ResponseWriter, r *http.Request) {
	// Database connection string
	dsn := "user=postgres password=postgres dbname=postgres port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	// Parse user
	var user Entity.User
	err = json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Validate UserType
	if user.UserType != 0 && user.UserType != 1 {
		http.Error(w, "Invalid user type", http.StatusBadRequest)
		return
	}

	// Set the registration date
	user.RegisterDate = time.Now()

	if user.UserType == 1 {
		user.StoreID = 0

		var lastAddressID int
		db.Table("loja_endereco").Select("MAX(id_endereco)").Scan(&lastAddressID)
		user.StoreOrAddress.AddressID = lastAddressID + 1

		err = db.Create(&user.StoreOrAddress).Error
		if err != nil {
			http.Error(w, "Failed to register user address", http.StatusInternalServerError)
			return
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
			http.Error(w, "Failed to register seller store", http.StatusInternalServerError)
			return
		}

		user.StoreID = user.StoreOrAddress.StoreID
		user.AddressID = user.StoreOrAddress.AddressID
	}

	err = db.Create(&user).Error
	if err != nil {
		http.Error(w, "Failed to register user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("User registered successfully"))
}
