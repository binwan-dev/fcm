package handlers

import (
	"time"

	"github.com/Atlantis-Org/fcm/models"
	"github.com/Atlantis-Org/fcm/utils"
)

func CreateApp(app *models.App) error {
	if app == nil {
		return utils.ErrParameterInvalid
	}

	var count int64
	utils.Db.Model(&app).Where("name = ?", app.Name).Count(&count)
	if count > 0 {
		return utils.ErrExisted
	}

	app.CreateAt = time.Now().Unix()

	return utils.Db.Model(&models.App{}).Create(app).Error
}

func CreateAppNamespace(namespace *models.AppNamespace) error {
	if namespace == nil {
		return utils.ErrParameterInvalid
	}

	var count int64
	utils.Db.Model(&namespace).Where("name = ? ", namespace.Name).Count(&count)
	if count > 0 {
		return utils.ErrExisted
	}

	namespace.CreateAt = time.Now().Unix()

	return utils.Db.Model(&models.AppNamespace{}).Create(namespace).Error
}

func CreateAppConfig(config *models.AppConfigInfo) error {
	if config == nil {
		return utils.ErrParameterInvalid
	}

	var count int64
	utils.Db.Model(&config).Where("name = ? ", config.Name).Count(&count)
	if count > 0 {
		return utils.ErrExisted
	}

	config.CreateAt = time.Now().Unix()

	return utils.Db.Model(&models.AppConfigInfo{}).Create(config).Error
}
