package Entity

type ProductFeatures struct {
	FeatureID int `gorm:"primaryKey;column:fk_caracteristica_id_caracteristica"`
	ProductID int `gorm:"primaryKey;column:fk_produto_id_produto"`
}

func (ProductFeatures) TableName() string {
	return "produto_caracteristica"
}
