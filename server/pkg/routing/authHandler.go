package routing

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/klopjq/telemedicine/config"
	"github.com/klopjq/telemedicine/internal/log"
)

func auth(cfg *config.Config, logger log.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		apiKey := c.GetHeader("X-API-KEY")
		if apiKey != cfg.ServerApiKey {
			logger.Error(http.StatusText(http.StatusUnauthorized))
			c.AbortWithStatusJSON(
				http.StatusUnauthorized,
				gin.H{
					"status": http.StatusText(http.StatusUnauthorized),
					"code":   http.StatusUnauthorized,
				},
			)
			return
		}
	}
}
