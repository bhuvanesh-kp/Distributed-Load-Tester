package main

import (
	"fmt"
	"time"
	"version_1/cmd/controller"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Version1 of Distribured-Load-Tester(DLT) ...")

	controller := controller.NewController()

	r := gin.Default()

	r.GET("/", controller.HealthChecker)

	r.Run() // application at localhost:8080/
	time.Sleep(time.Second * 10)
}
