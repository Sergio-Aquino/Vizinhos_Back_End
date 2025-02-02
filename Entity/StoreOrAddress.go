package Entity

type StoreOrAddress struct {
	StoreID     int    `gorm:"primaryKey;column:id_loja;uniqueIndex:store_address_idx" json:"StoreID"`
	AddressID   int    `gorm:"primaryKey;column:id_endereco;uniqueIndex:store_address_idx;autoIncrement" json:"AddressID"`
	CEP         string `gorm:"column:cep" json:"CEP"`
	Street      string `gorm:"column:logradouro" json:"Street"`
	Number      string `gorm:"column:numero" json:"Number"`
	Complement  string `gorm:"column:complemento" json:"Complement"`
	StoreName   string `gorm:"column:nome_loja" json:"StoreName"`
	Description string `gorm:"column:descricao_loja" json:"Description"`
}

func (StoreOrAddress) TableName() string {
	return "loja_endereco"
}
