package models

import "time"

type Chat struct {
	IDChat    uint      `gorm:"column:id_chat;primaryKey;autoIncrement" json:"id_chat"`
	IDUser    *uint     `gorm:"column:id_user" json:"id_user"`
	IDPenjual *uint     `gorm:"column:id_penjual" json:"id_penjual"`
	Chat      string    `gorm:"column:chat" json:"chat"`
	Sender    string    `gorm:"column:sender" json:"sender"`
	IsRead    bool      `gorm:"column:is_read;default:false" json:"is_read"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`

	User    *User    `gorm:"foreignKey:IDUser;references:IDUser" json:"user"`
	Penjual *Penjual `gorm:"foreignKey:IDPenjual;references:IDPenjual" json:"penjual"`
}

func (Chat) TableName() string {
	return "chat"
}
