package user

import (
	"fmt"

	"../../handler"
	"../../pkg/errno"

	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
)

// Create 创建一个新用户
func Create(c *gin.Context) {
	var req CreateRequest
	if err := c.Bind(&req); err != nil {
		SendResponse(c, errno.ErrBind, nil)
		return
	}
	admin := c.Param("user_name")
	log.Infof("URL param user_name: %s", admin)
	desc := c.Query("desc")
	log.Infof("URL key param desc: %s", desc)
	contentType := c.GetHeader("Content-Type")
	log.Infof("Header Content-Type: %s", contentType)
	log.Debugf("user_name is: [%s], password is [%s]", req.UserName, req.Password)
	
	if req.UserName == "" {
		SendResponse(c, errno.New(errno.ErrUserNotFound, fmt.Errorf("user_name can not found in db: xxx.xxx.xxx.xxx")), nil)
		return
	}
	if req.Password == "" {
		SendResponse(c, fmt.Errorf("password is empty"), nil)
	}

	resp := CreateResponse {
		UserName: req.UserName,
	}
	SendResponse(c, nil, resp)
}