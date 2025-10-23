package rest

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		clientIP := c.ClientIP()
		fmt.Printf("This is the clients IP: %s\n", clientIP)
		fmt.Println("Logging request...")
		method := c.Request.Method
		path := c.Request.URL.Path
		currentTime := time.Now().Format("2006-01-02 15:04:05")
		fmt.Printf("[LOG] %s | %s | %s | %s\n", clientIP, currentTime, method, path)
		c.Next()
		statusCode := c.Writer.Status()
		fmt.Printf("[LOG] Completed %s %s with %d\n", method, path, statusCode)
	}

}
