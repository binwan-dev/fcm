package models

type App struct {
	Id         int `gorm:"primaryKey;autoIncrement"`
	GroupId    int
	Name       string
	CreateAt   int64
	Namespaces []AppNamespace `gorm:"-"`
}
