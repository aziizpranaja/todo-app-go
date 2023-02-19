package models

type Todo struct {
	Id         int32  `gorm:"primaryKey" json:"id"`
	Title      string `gorm:"type:varchar(150)" binding:"required" json:"title"`
	Desc       string `gorm:"type:text" binding:"required" json:"desc"`
	IsComplete bool   `gorm:"default:false" binding:"required" json:"is_complete"`
	UserId     int32
	User       User
}