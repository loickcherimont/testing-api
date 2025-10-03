package main

import (
	"github.com/gin-gonic/gin"
	"github.com/loickcherimont/testing-api/controller"
)

func main() {
	router := gin.Default()

	router.GET("/ping", controller.PingTest)
	router.GET("/tests", controller.GetAllTests)
	router.GET("/tests/:id", controller.GetTestById)
	router.Run()
}
