package models

type Alamat struct {
	IDAlamat    uint   `gorm:"column:id_alamat;primaryKey;autoIncrement" json:"id_alamat"`
	IDUser      uint   `gorm:"column:id_user" json:"id_user"`
	Alamat      string `gorm:"column:alamat" json:"alamat"`
	TeleponA    string `gorm:"column:teleponA" json:"teleponA"`
	NamaA       string `gorm:"column:namaA" json:"namaA"`
	AlamatUtama bool   `gorm:"column:alamat_utama" json:"alamat_utama"`

	User User `gorm:"foreignKey:IDUser;references:IDUser"`
}

func (Alamat) TableName() string {
	return "alamat"
}
