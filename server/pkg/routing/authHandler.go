package routing

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/klopjq/telemedicine/internal/log"
)

func auth(logger log.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		apiKey := c.GetHeader("X-API-KEY")
		//@TODO: fix this to get from the users information
		if apiKey != "apiKey" {
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
