package user

import (
	"go/tiny_http_server/model"
	"go/tiny_http_server/pkg/errno"
	"go/tiny_http_server/util"

	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
	"github.com/lexkong/log/lager"
)

// Create Add new user to the database
// @Summary Add new user to the database
// @Description Add a new user to the database
// @Tags user
// @Accept  json
// @Produce  json
// @Param user body user.CreateRequest true "Create a new user"
// @Success 200 {object} user.CreateResponse "{"code":0,"msg":"Ok","data":{"user_name":"pyl"}}"
// @Router /user [post]
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
