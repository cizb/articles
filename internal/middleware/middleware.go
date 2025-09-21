package middleware

import (
	"github.com/gin-gonic/gin"
)

func CustomHeader() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("X-App-Version", "1.0.0")
		c.Next()
	}
}
