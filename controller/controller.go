package controller

import (
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/loickcherimont/testing-api/service"
)

// Test to verify if API runs
// todo - testing package
func PingTest(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

func GetAllTests(c *gin.Context) {
	tests := service.GetAllTests()
	c.JSON(http.StatusOK, tests)
}

func GetTestById(c *gin.Context) {
	stringId := c.Param("id")
	id, err := strconv.ParseUint(stringId, 10, 64)
	if err != nil {
		log.Printf("Error: %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	uid := uint(id)
	test := service.GetTestById(uid)

	if test.Status == "" && test.Title == "" && len(test.Details) == 0 {
		err = errors.New("sorry but this test doesn't exist")
		log.Printf("Error: %v\n", err)
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, test)
}
