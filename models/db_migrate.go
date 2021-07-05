package models

import "github.com/Atlantis-Org/fcm/utils"

func DbMigrate() {
	utils.Db.AutoMigrate(&Group{}, &App{}, &AppNamespace{}, &AppConfigInfo{})
}
