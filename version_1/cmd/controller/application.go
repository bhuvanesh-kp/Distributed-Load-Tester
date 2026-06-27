package controller

import (
	"github.com/gin-gonic/gin"
)

type appControllerInterface interface {
	HealthChecker(c *gin.Context)
}

func NewController() appControllerInterface {
	return &appStruct{}
}

type appStruct struct{}

func (ctl *appStruct) HealthChecker(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "Application is alive",
	})
}
