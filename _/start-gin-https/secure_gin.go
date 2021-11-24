// Integration: Gin
//  https://github.com/unrolled/secure#gin
package main

import (
	"github.com/gin-gonic/gin"
	"github.com/unrolled/secure" // or "gopkg.in/unrolled/secure.v1"
)

func main() {
	secureMiddleware := secure.New(secure.Options{
		FrameDeny: true,
	})
	secureFunc := func() gin.HandlerFunc {
		return func(c *gin.Context) {
			err := secureMiddleware.Process(c.Writer, c.Request)

			// If there was an error, do not continue.
			if err != nil {
				c.Abort()
				return
			}

			// Avoid header rewrite if response is a redirection.
			if status := c.Writer.Status(); status > 300 && status < 399 {
				c.Abort()
			}
		}
	}()

	router := gin.Default()
	router.Use(secureFunc)

	router.GET("/", func(c *gin.Context) {
		// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/X-Frame-Options
		c.String(200, "X-Frame-Options header is now `DENY`.")
	})

	router.RunTLS(":3000", "cert.pem", "key.pem")
}
