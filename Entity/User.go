package Entity

import (
	"time"
)

type User struct {
	CPF            string         `gorm:"primaryKey;column:cpf" json:"CPF"`
	UserType       int            `gorm:"column:usuario_tipo" json:"UserType"`
	StoreID        int            `gorm:"column:fk_id_loja" json:"StoreID"`
	AddressID      int            `gorm:"column:fk_id_endereco" json:"AddressID"`
	PhoneNumber    string         `gorm:"column:telefone" json:"PhoneNumber"`
	Email          string         `gorm:"column:email" json:"Email"`
	RegisterDate   time.Time      `gorm:"column:data_cadastro" json:"register_date"`
	StoreOrAddress StoreOrAddress `gorm:"foreignKey:StoreID,AddressID;references:StoreID,AddressID" json:"StoreOrAddress"`
}

func (User) TableName() string {
	return "usuario"
}
