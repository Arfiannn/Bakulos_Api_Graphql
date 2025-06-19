package models

type Keranjang struct {
	IDKeranjang uint   `gorm:"column:id_keranjang;primaryKey;autoIncrement" json:"id_keranjang"`
	IDProduct   uint   `gorm:"column:id_product" json:"id_product"`
	IDUser      uint   `gorm:"column:id_user" json:"id_user"`
	Jumlah      int    `gorm:"column:jumlah" json:"jumlah"`
	SizeK       string `gorm:"column:sizeK" json:"sizeK"`

	Product Product `gorm:"foreignKey:IDProduct" json:"product"`
	User    User    `gorm:"foreignKey:IDUser" json:"user"`
}
