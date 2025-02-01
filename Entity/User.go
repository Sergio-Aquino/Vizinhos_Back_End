package Entity

import (
	"time"
)

type User struct {
	CPF            string           `gorm:"primaryKey;column:cpf"`
	UserType       int              `gorm:"column:usuario_tipo"`
	StoreID        string           `gorm:"column:fk_id_loja"`
	AddressID      string           `gorm:"column:fk_id_endereco"`
	PhoneNumber    string           `gorm:"column:telefone"`
	Email          string           `gorm:"column:email"`
	RegisterDate   time.Time        `gorm:"column:data_cadastro"`
	StoreOrAddress []StoreOrAddress `gorm:"foreignKey:StoreID,AddressID;references:StoreID,AddressID"`
}

func (User) TableName() string {
	return "usuario"
}
