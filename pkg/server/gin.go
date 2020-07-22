package server

import (
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

var RequestIdHeader = "x-request-id"

func ginLogger(l *logrus.Logger) gin.HandlerFunc {
	hostname, err := os.Hostname()
	if err != nil {
		hostname = "unknow"
	}

	gin.Logger()

	return func(c *gin.Context) {
		path := c.Request.URL.Path
		start := time.Now()

		c.Next()

		stop := time.Now()
		latency := stop.Sub(start)
		statusCode := c.Writer.Status()
		clientIP := c.ClientIP()
		clientUserAgent := c.Request.UserAgent()
		referer := c.Request.Referer()

		dataLength := c.Writer.Size()
		if dataLength < 0 {
			dataLength = 0
		}

		//reqID := c.GetString(RequestIdHeader)
		reqID := c.GetHeader(RequestIdHeader)
		method := c.Request.Method

		entry := l.WithFields(logrus.Fields{
			"Hostname":    hostname,
			"Path":        path,
			"Latency":     latency,
			"Code":        statusCode,
			"IP":          clientIP,
			"User-Agent":  clientUserAgent,
			"Referer":     referer,
			"Data-Length": dataLength,
			"RequestID":   reqID,
			"Method":      method,
		})

		if len(c.Errors) > 0 {
			entry.Error(c.Errors.ByType(gin.ErrorTypePrivate).String())
		} else {
			msg := "HTTP Request"
			if statusCode >= 500 {
				entry.Error(msg)
			} else if statusCode >= 400 {
				entry.Warn(msg)
			} else {
				entry.Info(msg)
			}
		}
	}
}
