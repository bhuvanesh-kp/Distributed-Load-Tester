package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func main(){
	fmt.Println("Version1 of Distribured-Load-Tester(DLT) ...");

	r := gin.Default();

	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "Application is alive",
		});
	})

	r.Run() // application at localhost:8080/
	time.Sleep(time.Second * 10);
}