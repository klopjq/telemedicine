package tmis

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/klopjq/telemedicine/internal/log"
)

type form struct {
	BodyTemperature  float64 `form:"field1,omitempty"`
	HeartRate        string  `form:"field2,omitempty"`
	BloodPressureSys int     `form:"field3,omitempty"`
	BloodPressureDia int     `form:"field4,omitempty"`
	OxygenSaturation string  `form:"field5,omitempty"`
	RespiratoryRate  string  `form:"field6,omitempty"`
	FieldSeven       string  `form:"field7,omitempty"`
	FieldEight       string  `form:"field8,omitempty"`
	ReceivedAt       time.Time
}

func UpdateHandler(logger log.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		var f form
		if err := c.Bind(&f); err != nil {
			logger.Error(err)
			c.AbortWithStatusJSON(http.StatusBadRequest,
				gin.H{
					"status": http.StatusText(http.StatusBadRequest),
					"code":   http.StatusBadRequest,
					"error":  err,
				})
			return
		}
		f.ReceivedAt = time.Now()
		c.JSON(http.StatusOK,
			gin.H{
				"status":  http.StatusText(http.StatusOK),
				"code":    http.StatusOK,
				"payload": f,
			})
	}
}
