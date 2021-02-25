package user

import (
	"strconv"

	"go/tiny_http_server/model"
	"go/tiny_http_server/handler"
	"go/tiny_http_server/pkg/errno"

	"github.com/gin-gonic/gin"
)

// Delete 删除用户
func Delete(c *gin.Context) {
	userID, _ := strconv.Atoi(c.Param("id"))
	if err := model.DeleteUser(uint64(userID)); err != nil {
		SendResponse(c, errno.ErrDatabase, nil)
		return
	}

	SendResponse(c, nil, nil)
}