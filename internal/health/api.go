package health

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/klopjq/telemedicine/internal/log"
)

func HealthHandler(logger log.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK,
			gin.H{
				"status": http.StatusText(http.StatusOK),
				"code":   http.StatusOK,
			})

	}
}
