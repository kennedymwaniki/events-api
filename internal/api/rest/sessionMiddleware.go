package rest

import (
	"github.com/gin-gonic/gin"
)

func SessionMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		sessionID := c.GetHeader("Session-ID")
		if sessionID == "" {
			c.AbortWithStatusJSON(401, gin.H{
				"error": "Session-ID header is required",
			})
			return
		}
		c.Next()
	}
}