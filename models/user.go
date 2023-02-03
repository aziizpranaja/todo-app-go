package models

type User struct {
	Id       int32  `gorm:"primaryKey" json:"id"`
	Email    string `gorm:"type:varchar(150);unique" json:"email"`
	Password string `gorm:"type:varchar(300)" json:"password"`
}