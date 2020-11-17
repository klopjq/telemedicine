package tmis

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/klopjq/telemedicine/internal/log"
)

type form struct {
	bodyTemperature  float64 `form:"field1,omitempty"`
	heartRate        string  `form:"field2,omitempty"`
	bloodPressureSys int     `form:"field3,omitempty"`
	bloodPressureDia int     `form:"field4,omitempty"`
	oxygenSaturation string  `form:"field5,omitempty"`
	respiratoryRate  string  `form:"field6,omitempty"`
	fieldSeven       string  `form:"field7,omitempty"`
	fieldEight       string  `form:"field8,omitempty"`
	receivedAt       time.Time
}

func UpdateHandler(logger log.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		var f form
		if err := c.ShouldBind(&f); err != nil {
			logger.Error(err)
			c.AbortWithStatusJSON(http.StatusUnprocessableEntity,
				gin.H{
					"status": http.StatusText(http.StatusUnprocessableEntity),
					"code":   http.StatusUnprocessableEntity,
					"error":  err,
				})
			return
		}
		f.receivedAt = time.Now()
		c.JSON(http.StatusOK,
			gin.H{
				"status":  http.StatusText(http.StatusOK),
				"code":    http.StatusOK,
				"payload": f,
			})
	}
}
