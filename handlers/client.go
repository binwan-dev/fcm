package handlers

import (
	"github.com/Atlantis-Org/fcm/models"
	"github.com/Atlantis-Org/fcm/utils"
	"gorm.io/gorm"
)

func GetConfigForApp(appName string) (error, *models.App, []models.AppConfigInfo) {
	var app models.App
	err := utils.Db.Where(&models.App{Name: appName}).First(&app).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil, []models.AppConfigInfo{}
	}
	if err != nil {
		return err, nil, []models.AppConfigInfo{}
	}

	err = utils.Db.Where("app_id = ? or (role = ? and group_id = ?) or role = ?", app.Id, models.Internal, app.GroupId, models.Public).Find(&app.Namespaces).Error
	if err == gorm.ErrRecordNotFound {
		return nil, &app, []models.AppConfigInfo{}
	}
	if err != nil {
		return err, &app, []models.AppConfigInfo{}
	}

	namespaceIds := make([]int, len(app.Namespaces))
	for i, namespace := range app.Namespaces {
		namespaceIds[i] = namespace.Id
	}

	var configs []models.AppConfigInfo
	err = utils.Db.Where("namespace_id in ?", namespaceIds).Find(&configs).Error
	if err == gorm.ErrRecordNotFound {
		return nil, &app, []models.AppConfigInfo{}
	}
	if err != nil {
		return err, &app, []models.AppConfigInfo{}
	}

	return nil, &app, configs
}
