package middleware

import (
	"go/tiny_http_server/handler"
	"go/tiny_http_server/pkg/errno"
	"go/tiny_http_server/pkg/token"

	"github.com/gin-gonic/gin"
)

// AuthMiddleware 解析Token
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if _, err := token.ParseRequest(c); err != nil {
			handler.SendResponse(c, errno.ErrTokenInvalid, nil)
			c.Abort()
			return
		}
		c.Next()
	}
}
