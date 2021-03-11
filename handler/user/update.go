package user

import (
	"strconv"

	"go/tiny_http_server/handler"
	"go/tiny_http_server/model"
	"go/tiny_http_server/pkg/errno"
	"go/tiny_http_server/util"

	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
	"github.com/lexkong/log/lager"
)

// Update a user info by the user id
// @Summary Update a user info by the user id
// @Description Update a user info by the user id
// @Tags user
// @Accept  json
// @Produce  json
// @Param id path uint64 true "The user database id"
// @Param user body model.UserModel true "The user info"
// @Success 200 {object} handler.Response "{"code":0,"msg":"Ok","data":null}"
// @Router /user/{id} [put]
func Update(c *gin.Context) {
	log.Info("Update() called.", lager.Data{"X-Request-Id": util.GetReqID(c)})
	userID, _ := strconv.Atoi(c.Param("id"))

	var u model.UserModel
	if err := c.Bind(&u); err != nil {
		SendResponse(c, errno.ErrBind, nil)
		return
	}
	// 根据Id更新用户
	u.Id = uint64(userID)
	if err := u.Validate(); err != nil {
		SendResponse(c, errno.ErrValidation, nil)
		return
	}
	if err := u.Encrypt(); err != nil {
		SendResponse(c, errno.ErrEncrypt, nil)
		return
	}
	if err := u.Update(); err != nil {
		SendResponse(c, errno.ErrDatabase, nil)
		return
	}
	SendResponse(c, nil, nil)
}
