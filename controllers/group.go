package controllers

import (
	"net/http"

	"github.com/Atlantis-Org/fcm/handlers"
	"github.com/Atlantis-Org/fcm/models"
	"github.com/gin-gonic/gin"
)

func CreateGroup(c *gin.Context) {
	var group models.Group
	err := c.ShouldBind(&group)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = handlers.CreateGroup(&group)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	c.Status(http.StatusOK)
}
