package user

import (
	"go/tiny_http_server/service"
	"go/tiny_http_server/handler"
	"go/tiny_http_server/pkg/errno"

	"github.com/gin-gonic/gin"
)

// List 获取用户列表
func List(c *gin.Context)  {
	var req ListRequest
	if err := c.Bind(&req); err != nil {
		SendResponse(c, errno.ErrBind, nil)
		return
	}
	infos, count, err := service.ListUser(req.UserName, req.Offset, req.Limit)
	if err != nil {
		SendRespone(c, err, nil)
		return
	}

	SendResponse(c, nil, ListResponse {
		Count: count,
		UserList: infos,
	})
}