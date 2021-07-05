package models

type AppNamespace struct {
	Id       int `gorm:"primaryKey;autoIncrement"`
	AppId    int
	GroupId  int
	Name     string
	CreateAt int64
	Role     Role
	Configs  []AppConfigInfo `gorm:"-"`
}
