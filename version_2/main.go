package main

import (
	"fmt"
	"version_2/cmd/controller"
	"version_2/internal/middleware"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	fmt.Println("Version1 of Distribured-Load-Tester(DLT) ...")

	controller := controller.NewController()

	r := gin.Default()
	middleware.PrometheusInit()

	// prometheus metric endpoint
	r.GET("/metrics", gin.WrapH(promhttp.Handler()))

	r.Use(middleware.TrackMetrics())

	// health point to check is application is live
	r.GET("/", controller.HealthChecker)

	// endpoint to test the is the load for single worker
	r.POST("/testendpoint", controller.TestEndPoint)

	r.Run()
}
