package routing

import (
	"github.com/klopjq/telemedicine/internal/tmis"
	"net/http"
	"strings"

	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"github.com/klopjq/telemedicine/internal/health"
	"github.com/klopjq/telemedicine/internal/log"
)

func New(logger log.Logger) http.Handler {
	router := gin.New()
	router.Use(
		//ginlog.GinLog(logger),
		gin.Logger(),
		gin.Recovery(),
		// Handle compressions
		func(c *gin.Context) {
			ae := c.Request.Header.Get("Accept-Encoding")
			switch true {
			case strings.Contains(ae, "br"):
				c.Request.Header.Set("Accept-Encoding", "br")
			case strings.Contains(ae, "gzip"):
				c.Request.Header.Set("Accept-Encoding", "gzip")
			}
			c.Next()
		},
		gzip.Gzip(gzip.DefaultCompression, gzip.WithDecompressFn(gzip.DefaultDecompressHandle)),
		//brotli.Brotli(cbrotli.WriterOptions{Quality: 5}, brotli.WithDecompressFn(brotli.DefaultDecompressHandle)),
	)

	api := router.Group("/v1")
	api.GET("/health", health.HealthHandler(logger))
	api.POST("/update", auth(logger), tmis.UpdateHandler(logger))
	//api.POST("/login", auth(cfg, logger), auth.LoginHandler(db, logger))

	return router
}
