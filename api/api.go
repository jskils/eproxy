package api

import (
	"eproxy/model"
	"eproxy/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Run(c *gin.Context) {
	var request model.RunRequest
	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "parameter error"})
		return
	}
	data, err := service.Run(request.Amount)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": data})
}

func Change(c *gin.Context) {
	var request model.IdRequest
	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "parameter error"})
		return
	}
	data, err := service.Change(request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": data})
}

func Release(c *gin.Context) {
	var request model.IdRequest
	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "parameter error"})
		return
	}
	data, err := service.Release(request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": data})
}
