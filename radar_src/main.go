package main

import (
	"github.com/gin-contrib/logger"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/radar", logger.SetLogger(
		logger.SetLogger(),
	), func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hi hacker",
		})
	})
	err := r.Run(":7777")
	if err != nil {
		return
	}
}
