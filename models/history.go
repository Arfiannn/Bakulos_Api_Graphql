package models

type History struct {
	IDHistory uint   `gorm:"column:id_history;primaryKey;autoIncrement" json:"id_history"`
	IDProduct uint   `gorm:"column:id_product" json:"id_product"`
	IDUser    uint   `gorm:"column:id_user" json:"id_user"`
	Jumlah    int    `gorm:"column:jumlah" json:"jumlah"`
	SizeH     string `gorm:"column:sizeH" json:"sizeH"`

	Product Product `gorm:"foreignKey:IDProduct;references:IDProduct" json:"product"`
	User    User    `gorm:"foreignKey:IDUser;references:IDUser" json:"user"`
}

func (History) TableName() string {
	return "history"
}
