package controllers

import (
	"net/http"
	"strconv"

	"github.com/Atlantis-Org/fcm/handlers"
	"github.com/Atlantis-Org/fcm/models"
	"github.com/gin-gonic/gin"
)

func GetAppForId(c *gin.Context) {
	appId, _ := strconv.Atoi(c.Param("appId"))
	err, app := handlers.GetAppForId(appId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, app)
}

func GetAppPaged(c *gin.Context) {
	pageNumber, _ := strconv.Atoi(c.Query("pageNumber"))
	pageSize, _ := strconv.Atoi(c.Query("pageSize"))
	err, paged := handlers.GetAppPages(pageNumber, pageSize)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, paged)
}

func CreateApp(c *gin.Context) {
	var app models.App
	err := c.ShouldBind(&app)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = handlers.CreateApp(&app)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	c.Status(http.StatusOK)
}

func CreateAppNamespace(c *gin.Context) {
	var namespace models.AppNamespace
	err := c.ShouldBindJSON(&namespace)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = handlers.CreateAppNamespace(&namespace)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	c.Status(http.StatusOK)
}

func GetAppNamespacePaged(c *gin.Context) {
	pageNumber, _ := strconv.Atoi(c.Query("pageNumber"))
	pageSize, _ := strconv.Atoi(c.Query("pageSize"))
	projectId, _ := strconv.Atoi(c.Query("projectId"))
	err, paged := handlers.GetAppNamespacePages(pageNumber, pageSize, projectId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, paged)
}

func GetAppConfigs(c *gin.Context) {
	namespaceId, _ := strconv.Atoi(c.Query("namespaceId"))
	err, configs := handlers.GetAppConfigs(namespaceId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, configs)
}

func CreateAppConfig(c *gin.Context) {
	var config models.AppConfigInfo
	err := c.ShouldBindJSON(&config)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = handlers.CreateAppConfig(&config)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	c.Status(http.StatusOK)
}

func ModifyAppConfig(c *gin.Context) {
	var config models.AppConfigInfo
	err := c.ShouldBindJSON(&config)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = handlers.ModifyAppConfig(&config)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	c.Status(http.StatusOK)
}
