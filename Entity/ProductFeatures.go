package Entity

type ProductFeatures struct {
	FeatureID string `gorm:"primaryKey;column:fk_caracteristica_id_caracteristica"`
	ProductID string `gorm:"primaryKey;column:fk_produto_id_produto"`
}

func (ProductFeatures) TableName() string {
	return "produto_caracteristica"
}
