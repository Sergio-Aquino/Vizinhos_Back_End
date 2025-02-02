package Entity

type Category struct {
	CategoryID  int    `gorm:"primaryKey;column:id_categoria"`
	Description string `gorm:"column:descricao"`
}

func (Category) TableName() string {
	return "categoria"
}
