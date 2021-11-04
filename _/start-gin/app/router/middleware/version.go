package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/ikuokuo/start-golang/_/start-gin/app/config"
)

// VersionMiddleware: add version to header
func VersionMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("X-START-GIN-VERSION", config.Version)
		c.Next()
	}
}
