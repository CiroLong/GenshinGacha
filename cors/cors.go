package Cors

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Options struct {
	Origin string
}

func CORS(options Options) gin.HandlerFunc {
	return func(c *gin.Context) {
		if options.Origin != "" {
			c.Writer.Header().Set("Access-Control-Allow-Origin", options.Origin)
		}
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "*")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		c.Writer.Header().Set("Access-Control-Expose-Headers", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		} else {
			c.Next()
		}
	}
}
