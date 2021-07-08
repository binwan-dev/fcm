package handlers

import (
	"time"

	"github.com/Atlantis-Org/fcm/models"
	"github.com/Atlantis-Org/fcm/utils"
	"gorm.io/gorm"
)

func GetAppPages(pageNumber, pageSize int) (error, utils.Paged) {
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

	var apps []models.App
	offest := (paged.PageNumber - 1) * 20
	err := utils.Db.Model(&models.App{}).Order("create_at desc").Limit(pageSize).Offset(offest).Find(&apps).Error
	paged.List = apps
	if err == gorm.ErrRecordNotFound {
		return nil, paged
	}
	if err != nil {
		return err, paged
	}

	err = utils.Db.Model(&models.App{}).Count(&paged.TotalRow).Error
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
