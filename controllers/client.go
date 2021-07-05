package controllers

import (
	"net/http"

	"github.com/Atlantis-Org/fcm/handlers"
	"github.com/Atlantis-Org/fcm/models"
	"github.com/gin-gonic/gin"
)

func GetClientConfig(c *gin.Context) {
	appName := c.Query("appname")
	err, app, configs := handlers.GetConfigForApp(appName)
	if err != nil {
		c.String(http.StatusInternalServerError, "Server Error")
		return
	}

	c.JSON(http.StatusOK,
		struct {
			App     models.App
			Configs []models.AppConfigInfo
		}{
			App:     *app,
			Configs: configs,
		})
}
