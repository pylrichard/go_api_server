package util

import (
	"github.com/gin-gonic/gin"
	"github.com/teris-io/shortid"
)

// GenShortID 生成短Id
func GenShortID() (string, error) {
	return shortid.Generate()
}

// GetReqID 获取请求Id
func GetReqID(c *gin.Context) string {
	v, ok := c.Get("X-Request-Id")
	if !ok {
		return ""
	}
	if reqID, ok := v.(string); ok {
		return reqID
	}

	return ""
}
