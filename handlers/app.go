package handlers

import (
	"time"

	"github.com/Atlantis-Org/fcm/models"
	"github.com/Atlantis-Org/fcm/utils"
	"gorm.io/gorm"
)

func GetAppForId(appId int) (error, *models.App) {
	var app models.App
	err := utils.Db.Find(&app, appId).Error
	if err != nil && err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	if err != nil {
		return err, nil
	}
	return nil, &app
}

func GetAppPages(pageNumber, pageSize int) (error, utils.Paged) {
	var apps []models.App
	db := utils.Db.Model(&models.App{}).Order("create_at desc")
	return toPaged(pageNumber, pageSize, db, &apps)
}

func GetAppNamespacePages(pageNumber, pageSize, projectId int) (error, utils.Paged) {
	var namespaces []models.AppNamespace
	db := utils.Db.Model(&models.AppNamespace{}).Order("create_at desc")
	return toPaged(pageNumber, pageSize, db, &namespaces)
}

func GetAppConfigPages(pageNumber, pageSize, namespaceId int) (error, utils.Paged) {
	var appConfigs []models.AppConfigInfo
	db := utils.Db.Model(&models.AppConfigInfo{}).Order("create_at desc")
	return toPaged(pageNumber, pageSize, db, &appConfigs)
}

func toPaged(pageNumber, pageSize int, db *gorm.DB, data interface{}) (error, utils.Paged) {
	if pageNumber <= 0 {
		pageNumber = 1
	}
	if pageSize <= 0 {
		pageSize = 20
	}

	paged := utils.Paged{
		PageNumber: pageNumber,
		PageSize:   pageSize,
	}

	offest := (paged.PageNumber - 1) * 20
	err := db.Limit(pageSize).Offset(offest).Find(data).Error
	paged.List = data
	if err == gorm.ErrRecordNotFound {
		return nil, paged
	}
	if err != nil {
		return err, paged
	}

	err = db.Count(&paged.TotalRow).Error
	if err == gorm.ErrRecordNotFound {
		return nil, paged
	}
	if err != nil {
		return err, paged
	}
	if paged.TotalRow == 0 {
		return nil, paged
	}

	val := int(paged.TotalRow) % paged.PageSize
	paged.TotalPage = int(paged.TotalRow) / paged.PageSize
	if val != 0 {
		paged.TotalPage += 1
	}
	return nil, paged

}

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
