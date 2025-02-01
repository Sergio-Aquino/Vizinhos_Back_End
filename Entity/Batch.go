package Entity

import (
	"time"
)

type Batch struct {
	BatchID           string    `gorm:"primaryKey;column:id_lote"`
	ProductID         string    `gorm:"column:fk_Produto_id_produto"`
	ManufacturingDate time.Time `gorm:"column:dt_fabricacao"`
	Discount          float64   `gorm:"column:valor_venda_desc"`
	Quantity          int       `gorm:"column:quantidade"`
}

func (Batch) TableName() string {
	return "produto_lote"
}
