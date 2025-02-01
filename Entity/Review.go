package Entity

type Review struct {
	UserCPF   string `gorm:"primaryKey;column:fk_Usuario_cpf"`
	StoreID   int    `gorm:"primaryKey;column:fk_id_Loja"`
	AddressID int    `gorm:"primaryKey;column:fk_id_Endereco"`
	Rating    int    `gorm:"column:avaliacao"`
	Comment   string `gorm:"column:comentario"`
}

func (Review) TableName() string {
	return "avaliacao"
}
