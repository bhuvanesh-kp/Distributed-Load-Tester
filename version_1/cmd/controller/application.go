package controller

import (
	"net/http"
	"version_1/internal/domain"

	"github.com/gin-gonic/gin"
)

type appControllerInterface interface {
	HealthChecker(c *gin.Context)
	TestEndPoint(c *gin.Context)
}

func NewController() appControllerInterface {
	return &appStruct{
		previousUrl:         make(map[string]any, 5),
		totalRequestHandled: 0,
	}
}

type appStruct struct {
	previousUrl         map[string]any
	totalRequestHandled int16
}

func (ctl *appStruct) HealthChecker(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Application is alive",
	})
}

func (ctl *appStruct) TestEndPoint(c *gin.Context) {
	var req domain.Endpoint

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": err.Error(),
		})
		return
	}

	// TODO: add worker logic to hit the provided endpoint

	c.JSON(http.StatusAccepted, gin.H{
		"status":      200,
		"url":         req.Url,
		"results":     ctl.previousUrl[req.Url],
		"cacheStatus": 0,
	})
}
