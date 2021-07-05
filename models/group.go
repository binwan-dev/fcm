package models

type Group struct {
	Id       int `gorm:"primaryKey;autoIncrement"`
	Name     string
	CreateAt int64
}
