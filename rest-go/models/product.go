package models

type Product struct {
	Id          uint   `json:"id" gorm:"primary_key"`
	Title       string `json:"title" gorm:"type:varchar(300)"`
	Dexcription string `json:"description" gorm:"type:varchar(300)"`
}
