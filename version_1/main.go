package main

import (
	"fmt"
	"version_1/cmd/controller"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Version1 of Distribured-Load-Tester(DLT) ...")

	controller := controller.NewController()

	r := gin.Default()

	r.GET("/", controller.HealthChecker)
	r.POST("/testendpoint", controller.TestEndPoint)

	r.Run()
}
