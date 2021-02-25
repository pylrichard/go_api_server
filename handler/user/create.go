package user

import (
	"fmt"

	"go/tiny_http_server/handler"
	"go/tiny_http_server/model"
	"go/tiny_http_server/pkg/errno"
	"go/tiny_http_server/util"

	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
	"github.com/lexkong/log/lager"
)

// Create 创建一个新用户
func Create(c *gin.Context) {
	log.Info("User Create() called.", lager.Data{"X-Request-Id": util.GetReqID(c)})
	var req CreateRequest
	if err := c.Bind(&req); err != nil {
		SendResponse(c, errno.ErrBind, nil)
		return
	}

	u := model.UserModel{
		UserName: req.UserName,
		Password: req.Password,
	}
	if err := u.Validate(); err != nil {
		SendResponse(c, errno.ErrValidation, nil)
		return
	}
	if err := u.Encrypt(); err != nil {
		SendResponse(c, errno.ErrEncrypt, nil)
		return
	}
	if err := u.Create(); err != nil {
		SendResponse(c, errno.ErrDatabase, nil)
		return
	}

	resp := CreateResponse{
		UserName: req.UserName,
	}
	SendResponse(c, nil, resp)
}
