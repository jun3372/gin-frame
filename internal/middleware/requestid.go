package middleware

import (
	"github.com/gin-gonic/gin"

	"github.com/jun3372/gin-frame/pkg/util"
)

const (
	// XRequestID 全局唯一ID key
	XRequestID = "X-Request-ID"
)

// RequestID 透传Request-ID，如果没有则生成一个
func RequestID() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Check for incoming header, use it if exists
		requestID := c.Request.Header.Get(XRequestID)

		// Create request id with UUID4
		if requestID == "" {
			requestID = util.GenUUID()
		}

		// Expose it for use in the application
		c.Set(XRequestID, requestID)

		// Set X-Request-ID header
		c.Writer.Header().Set(XRequestID, requestID)
		c.Next()
	}
}
