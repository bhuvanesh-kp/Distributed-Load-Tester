package controller

import (
	"net/http"
	"version_1/internal/domain"
	"version_1/internal/service"

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
	srv                 service.NewServiceStruct
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

	_, ok := ctl.previousUrl[req.Url]

	if !ok {
		res := ctl.srv.LoadWithSingleWorker(req.Url)
		ctl.previousUrl[req.Url] = res
	}

	c.JSON(http.StatusAccepted, gin.H{
		"status":      200,
		"results":     ctl.previousUrl[req.Url],
		"cacheStatus": 0,
	})
}
