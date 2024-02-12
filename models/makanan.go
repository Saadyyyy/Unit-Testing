package models

type Makanan struct {
	Id   uint   `json:"id" gorm:"primarykey"`
	Nama string `json:"nama" gorm:"type:varchar(255)"`
}
