package api

import (
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (a *API) zapLogging() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		latency := time.Since(start)
		reqMethod := c.Request.Method
		reqUri := c.Request.RequestURI
		statusCode := c.Writer.Status()

		a.logger.Info("Request",
			zap.Int("status", statusCode),
			zap.String("method", reqMethod),
			zap.String("uri", reqUri),
			zap.Duration("latency", latency),
		)
		c.Next()
	}
}
