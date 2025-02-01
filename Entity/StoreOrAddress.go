package Entity

type StoreOrAddress struct {
	StoreID     string `gorm:"primaryKey;column:id_loja"`
	AddressID   string `gorm:"primaryKey;column:id_endereco"`
	CEP         string `gorm:"column:cep"`
	Street      string `gorm:"column:logradouro"`
	Number      string `gorm:"column:numero"`
	Complement  string `gorm:"column:complemento"`
	StoreName   string `gorm:"column:nome_loja"`
	Description string `gorm:"column:descricao_loja"`
}

func (StoreOrAddress) TableName() string {
	return "loja_endereco"
}
