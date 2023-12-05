package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/radar", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hi ",
		})
	})
	err := r.Run(":7777")
	if err != nil {
		return
	}
}
