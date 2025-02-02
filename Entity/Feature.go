package Entity

type Feature struct {
	FeatureID   int    `gorm:"primaryKey;column:id_caracteristica"`
	Description string `gorm:"column:descricao"`
}

func (Feature) TableName() string {
	return "caracteristica"
}
