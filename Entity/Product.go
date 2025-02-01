package Entity

type Product struct {
	ProductID          string  `gorm:"primaryKey;column:id_produto"`
	StoreID            string  `gorm:"column:fk_loja_endereco_id_loja"`
	AddressID          string  `gorm:"column:fk_loja_endereco_id_endereco"`
	CategoryID         string  `gorm:"column:fk_categoria_id_categoria"`
	DaysToExp          int     `gorm:"column:dias_vcto"`
	SellPrice          float64 `gorm:"column:valor_venda"`
	ManufacturingPrice float64 `gorm:"column:valor_custo"`
	Size               string  `gorm:"column:tamanho"`
	Description        string  `gorm:"column:descricao"`
}

func (Product) TableName() string {
	return "produto"
}
