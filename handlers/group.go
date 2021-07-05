package handlers

import (
	"time"

	"github.com/Atlantis-Org/fcm/models"
	"github.com/Atlantis-Org/fcm/utils"
)

func CreateGroup(group *models.Group) error {
	if group == nil {
		return utils.ErrParameterInvalid
	}

	var count int64
	utils.Db.Model(&group).Where("name = ?", group.Name).Count(&count)
	if count > 0 {
		return utils.ErrExisted
	}

	group.CreateAt = time.Now().Unix()

	return utils.Db.Model(&models.Group{}).Create(group).Error
}
