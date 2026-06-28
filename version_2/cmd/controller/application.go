package controller

import (
	"net/http"
	"time"
	"version_2/internal/domain"
	"version_2/internal/service"

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
		CacheHits:           0,
		srv:                 service.NewService(),
	}
}

type appStruct struct {
	previousUrl         map[string]any
	totalRequestHandled int16
	srv                 service.AppServiceInterface
	CacheHits           int16
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
	startTimer := time.Now()

	_, ok := ctl.previousUrl[req.Url]

	if !ok {
		res := ctl.srv.LoadWithSingleWorker(req.Url)
		ctl.previousUrl[req.Url] = res
		ctl.totalRequestHandled++
	} else {
		ctl.CacheHits++
	}

	duration := time.Since(startTimer)

	c.JSON(http.StatusAccepted, gin.H{
		"status":                    200,
		"results":                   ctl.previousUrl[req.Url],
		"CacheHits":                 ctl.CacheHits,
		"TotalUniqueRequestHandled": ctl.totalRequestHandled,
		"TimeTakenByWorker":         time.Duration(duration.Microseconds()),
	})
}
