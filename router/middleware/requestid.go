package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/satori/go.uuid"
)

// RequestID 从请求头中获取X-Request-Id
func RequestID() gin.HandlerFunc {
	return func(g *gin.Context) {
		// Check for incoming header, use it if exist
		requestID := c.Request.Header.Get("X-Request-Id")
		// Create request id with UUID4
		if requestID == "" {
			u4, _ := uuid.NewV4()
			requestID = u4.String()
		}
		// Expose it for use in the application
		c.Set("X-Request-Id", requestID)
		// Set X-Request-Id header
		c.Writer.Header().Set("X-Request-Id", requestID)
		c.Next()
	}
}