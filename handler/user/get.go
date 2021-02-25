package user

import (
	"go/tiny_http_server/handler"
	"go/tiny_http_server/model"
	"go/tiny_http_server/pkg/errno"

	"github.com/gin-gonic/gin"
)

// Get 获取用户
func Get(c *gin.Context) {
	userName := c.Param("user_name")
	user, err := model.GetUser(userName)
	if err != nil {
		SendResponse(c, errno.ErrUserNotFound, nil)
		return
	}
	SendResponse(c, nil, user)
}