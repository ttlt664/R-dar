package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	retalog "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"log"
	"math"
	"os"
	"time"
)

func Logger() gin.HandlerFunc {
	// 开启一个文件
	logPath := "log/request"
	src, err := os.OpenFile(logPath, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		fmt.Println("err:", err)
	}

	// 将日志输出给上面的目录
	logWriter, _ := retalog.New(
		logPath+"%Y-%m-%d.log",
		retalog.WithMaxAge(7*24*time.Hour),
		retalog.WithRotationTime(24*time.Hour),
	)

	// 将报错等级直接输出给文件.log
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

	// 创建日志记录器
	logger := log.New(src, "", 0)
	logger.SetOutput(logWriter)

	// 使用中间件的处理函数
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

		// 记录访问信息
		logEntry := fmt.Sprintf(
			"HostName: %s, Status: %d, SpendTime: %s, IP: %s, Method: %s, Path: %s, DataSize: %d, Agent: %s",
			hostName, statusCode, spendTime, clientIp, method, path, dataSize, userAgent,
		)

		// 输出日志
		logger.Println(logEntry)

		// 根据状态码输出不同级别的日志
		if statusCode >= 500 {
			log.Fatal(logEntry)
		} else if statusCode >= 400 {
			log.Println(logEntry)
		} else {
			log.Println(logEntry)
		}
	}
}
