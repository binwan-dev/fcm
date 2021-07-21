package models

type AppConfigInfo struct {
	Id          int `gorm:"primaryKey;autoIncrement"`
	NamespaceId int
	Name        string
	Data        string
	CreateAt    int64
	UpdateAt    int64
	ValidType   ValidType
	Sort        int16
}
