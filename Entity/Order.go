package Entity

import (
	"time"
)

type Order struct {
	OrderID    string    `gorm:"primaryKey;column:id_pedido"`
	UserCPF    string    `gorm:"column:fk_usuario_cpf"`
	BatchID    string    `gorm:"column:fk_lote_id_lote"`
	Price      float64   `gorm:"column:valor"`
	Quantity   int       `gorm:"column:quantidade"`
	Date       time.Time `gorm:"column:data_pedido"`
	Status     string    `gorm:"column:status_pedido"`
	LastUpdate string    `gorm:"column:hora_atualizacao"`
}

func (Order) TableName() string {
	return "pedido"
}
