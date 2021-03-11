package user

import (
	"go/tiny_http_server/service"
	"go/tiny_http_server/handler"
	"go/tiny_http_server/pkg/errno"

	"github.com/gin-gonic/gin"
)

// List users in the database
// @Summary List users in the database
// @Description List users in the database
// @Tags user
// @Accept  json
// @Produce  json
// @Param user body user.ListRequest true "List Users"
// @Success 200 {object} user.ListResponse "{"code":0,"msg":"OK","data":{"Count":1,"UserList":[{"id":0,"user_name":"admin","random":"user 'admin' get random string 'EnqntiSig'","pwd":"$2a$10$veGcArz47VGj7l9xN7g2iuT9TF21jLI1YGXarGzvARNdnt4inC9PG","created_time":"2018-05-28 00:25:33","updated_time":"2018-05-28 00:25:33"}]}}"
// @Router /user [get]
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