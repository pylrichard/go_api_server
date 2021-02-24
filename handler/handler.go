package handler

import (
	"net/http"

	"../pkg/errno"

	"github.com/gin-gonic/gin"
)

// Response Code和Msg通过DecodeErr()解析error类型变量得来
type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

// SendResponse 发送响应，供所有服务模块返回时调用
func SendResponse(c *gin.Context, err error, data interface{}) {
	code, msg := errno.DecodeErr(err)
	c.JSON(http.StatusOK, Response {
		Code:	code,
		Msg:	msg,
		Data:	data,
	})
}
