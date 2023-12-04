package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	retalog "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	log "github.com/sirupsen/logrus"
	"math"
	"os"
	"time"
)

func Logger() gin.HandlerFunc {
	logPath := "log/request"
	src, err := os.OpenFile(logPath, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		fmt.Println("open logfile err:", err)
	}

	logWriter, _ := retalog.New(
		logPath+"%Y-%m-%d.log",
		retalog.WithMaxAge(7*24*time.Hour),
		retalog.WithRotationTime(24*time.Hour),
	)

	writeMap := lfshook.WriterMap{
		lfshook.InfoLevel:  logWriter,
		lfshook.FatalLevel: logWriter,
		lfshook.WarnLevel:  logWriter,
		lfshook.ErrorLevel: logWriter,
		lfshook.DebugLevel: logWriter,
		lfshook.PanicLevel: logWriter,
	}
	hook := lfshook.NewHook(writeMap, &lfshook.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})

	logger := log.New(src, "", 0)
	logger.SetOutput(logWriter)

	return func(c *gin.Context) {
		startTime := time.Now()
		c.Next()
		stopTime := time.Since(startTime)
		spendTime := fmt.Sprintf("%d ms", int(math.Ceil(float64(stopTime.Nanoseconds())/1000000.0)))
		hostName, err := os.Hostname()
		if err != nil {
			hostName = "unknown"
		}
		statusCode := c.Writer.Status()
		clientIp := c.ClientIP()
		userAgent := c.Request.UserAgent()
		dataSize := c.Writer.Size()
		if dataSize < 0 {
			dataSize = 0
		}
		method := c.Request.Method
		path := c.Request.RequestURI

		logEntry := fmt.Sprintf(
			"HostName: %s, Status: %d, SpendTime: %s, IP: %s, Method: %s, Path: %s, DataSize: %d, Agent: %s",
			hostName, statusCode, spendTime, clientIp, method, path, dataSize, userAgent,
		)

		logger.Println(logEntry)

		if statusCode >= 500 {
			log.Fatal(logEntry)
		} else if statusCode >= 400 {
			log.Println(logEntry)
		} else {
			log.Println(logEntry)
		}
	}
}
