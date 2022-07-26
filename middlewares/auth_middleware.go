package middlewares

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/raj23manj/demo-app-golang/log/zap_uber"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()

		// Set example variable
		c.Set("example", "12345")

		// before request

		c.Next()

		// after request
		latency := time.Since(t)
		log.Print("latency")
		log.Print(latency)

		zap_uber.Info("middleware called after request",
			zap_uber.Field("latency", latency))

		// access the status we are sending
		status := c.Writer.Status()
		log.Println(status)
	}
}
