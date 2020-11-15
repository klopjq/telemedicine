package ginlog

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/klopjq/telemedicine/internal/log"
)

func GinLog(logger log.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		query := c.Request.URL.RawQuery
		ctx := c.Request.Context()
		ctx = log.WithRequest(ctx, c.Request)
		c.Request = c.Request.WithContext(ctx)
		c.Next()

		logger.With(ctx,
			"ip", c.ClientIP(),
			"ua", c.Request.UserAgent(),
			"method", c.Request.Method,
			"status", c.Writer.Status(),
			"path", path,
			"query", query,
			"latency", time.Now().Sub(start).Microseconds(),
			"length", c.Writer.Size(),
		).Info()
	}
}
