package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
)

func Logger() gin.HandlerFunc {
	logPath := "log/"
	_, err := os.OpenFile(logPath, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		fmt.Println("open logfile err:", err)
	}

	return nil
}
