package user

import (
	"strconv"

	"go/tiny_http_server/model"
	"go/tiny_http_server/handler"
	"go/tiny_http_server/pkg/errno"

	"github.com/gin-gonic/gin"
)

// Delete user by the user id
// @Summary Delete user by the user id
// @Description Delete user by the user id
// @Tags user
// @Accept  json
// @Produce  json
// @Param id path uint64 true "The user database id"
// @Success 200 {object} handler.Response "{"code":0,"msg":"Ok","data":null}"
// @Router /user/{id} [delete]
func Delete(c *gin.Context) {
	userID, _ := strconv.Atoi(c.Param("id"))
	if err := model.DeleteUser(uint64(userID)); err != nil {
		SendResponse(c, errno.ErrDatabase, nil)
		return
	}

	SendResponse(c, nil, nil)
}