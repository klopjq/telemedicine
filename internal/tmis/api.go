package tmis

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/klopjq/telemedicine/internal/log"
)

type form struct {
	BodyTemperature  string `form:"field1,omitempty"`
	HeartRate        string `form:"field2,omitempty"`
	BloodPressureSys string `form:"field3,omitempty"`
	BloodPressureDia string `form:"field4,omitempty"`
	OxygenSaturation string `form:"field5,omitempty"`
	RespiratoryRate  string `form:"field6,omitempty"`
	FieldSeven       string `form:"field7,omitempty"`
	FieldEight       string `form:"field8,omitempty"`
}

func UpdateHandler(logger log.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		var f form
		if err := c.ShouldBind(&f); err != nil {
			c.AbortWithStatusJSON(http.StatusUnprocessableEntity,
				gin.H{
					"status": http.StatusText(http.StatusUnprocessableEntity),
					"code":   http.StatusUnprocessableEntity,
					"error":  err,
				})
			return
		}
		c.JSON(http.StatusOK,
			gin.H{
				"status":  http.StatusText(http.StatusOK),
				"code":    http.StatusOK,
				"payload": f,
			})
	}
}
