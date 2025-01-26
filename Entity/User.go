package Entity

import "time"

type User struct {
	CPF            string
	UserType       int
	StoreOrAddress []StoreOrAddress
	PhoneNumber    string
	Email          string
	RegisterDate   time.Time
}
