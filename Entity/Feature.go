package Entity

type Feature struct {
	FeatureID   string `gorm:"primaryKey;column:id_caracteristica"`
	Description string `gorm:"column:descricao"`
}

func (Feature) TableName() string {
	return "caracteristica"
}
