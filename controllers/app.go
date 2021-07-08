package controllers

import (
	"net/http"
	"strconv"

	"github.com/Atlantis-Org/fcm/handlers"
	"github.com/Atlantis-Org/fcm/models"
	"github.com/gin-gonic/gin"
)

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
