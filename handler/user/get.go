package user

import (
	"go/tiny_http_server/handler"
	"go/tiny_http_server/model"
	"go/tiny_http_server/pkg/errno"

	"github.com/gin-gonic/gin"
)

// Get an user by the user name
// @Summary Get an user by the user name
// @Description Get an user by the user name
// @Tags user
// @Accept  json
// @Produce  json
// @Param user_name path string true "UserName"
// @Success 200 {object} model.UserModel "{"code":0,"msg":"Ok","data":{"user_name":"pyl","pwd":"$2a$10$E0kwtmtLZbwW/bDQ8qI8e.eHPqhQOW9tvjwpyo/p05f/f4Qvr3OmS"}}"
// @Router /user/{user_name} [get]
func Get(c *gin.Context) {
	userName := c.Param("user_name")
	user, err := model.GetUser(userName)
	if err != nil {
		SendResponse(c, errno.ErrUserNotFound, nil)
		return
	}
	SendResponse(c, nil, user)
}