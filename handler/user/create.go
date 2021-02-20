package user

import (
	"fmt"
	"net/http"

	"../../pkg/errno"

	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
)

// Create 创建一个新用户
func Create(c *gin.Context) {
	var req struct {
		UserName string `json:"user_name`
		Password string `json:"password"`
	}

	var err error
	if err := c.Bind(&req); err != nil {
		c.JSON(http.StatusOK, gin.H{"error": errno.ErrBind})
		return
	}
	log.Debugf("user_name is: [%s], password is [%s]", req.UserName, req.Password)
	if req.UserName == "" {
		err = errno.New(errno.ErrUserNotFound, fmt.Errorf("user_name can not found in db: xxx.xxx.xxx.xxx")).Add("This is add msg.")
		log.Errorf(err, "Get an error")
	}
	if errno.IsErrUserNotFound(err) {
		log.Debug("err type is ErrUserNotFound")
	}
	if req.Password == "" {
		err = fmt.Errorf("password is empty")
	}
	code, msg := errno.DecodeErr(err)
	c.JSON(http.StatusOK, gin.H{"code": code, "msg": msg})
}